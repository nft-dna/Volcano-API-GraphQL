// Package svc implements monitoring and scanning services of the API server.
package svc

import (
	"artion-api-graphql/internal/types"
	"time"

	"github.com/ethereum/go-ethereum/common"
	eth "github.com/ethereum/go-ethereum/core/types"
)

func newTokenContract(evt *eth.Log, lo *logObserver) {
	// MM todo ERC20 TokenCreated event

	if len(evt.Data) != 64 || len(evt.Topics) != 1 {
		log.Errorf("not Factory::TokenCreated() event #%d/#%d; expected 64 bytes of data, %d given; expected 1 topic, %d given",
			evt.BlockNumber, evt.Index, len(evt.Data), len(evt.Topics))
		return
	}

	// make meme token address
	ca := common.Address{}
	ca.SetBytes(evt.Data[32:64])
	meme := types.Collection{
		Address:      ca,
		IsActive:     true,
		IsOurFactory: true,
	}
	log.Infof("found (newTokenContract) MEME contract %s", ca.String())

	if err := extendMemeTokenDetails(&meme, evt, lo); err != nil {
		log.Criticalf("failed to load meme collection %s; %s", meme.Address.String(), err.Error())
		return
	}

	if err := repo.StoreMemeToken(&meme); err != nil {
		log.Criticalf("can not store meme collection %s; %s", meme.Address.String(), err.Error())
		return
	}

	// add observed contract based on the collection
	addObservedContract(&meme, evt)
	log.Infof("new meme collection %s found at %s", meme.Name, meme.Address.String())
}

func extendMemeTokenDetails(meme *types.Collection, evt *eth.Log, _ *logObserver) (err error) {
	// NFT contract type is derived from the factory contract type
	/*nft.Type, err = repo.NFTContractType(&evt.Address)
	if err != nil {
		log.Errorf("contract %s type not known; %s", evt.Address.String(), err.Error())
		return err
	}
	log.Debugf("NFT contract %s is %s", nft.Address.String(), nft.Type)
	*/
	meme.Type = types.ContractTypeERC20

	blk, err := repo.GetHeader(evt.BlockNumber)
	if err != nil {
		log.Errorf("header #%d not available; %s", evt.BlockNumber, err.Error())
		return err
	}
	meme.Created = types.Time(time.Unix(int64(blk.Time), 0))

	return extendMemeTokenMetadata(meme)
}

func extendMemeTokenMetadata(meme *types.Collection) (err error) {
	meme.Name, err = repo.MemeTokenName(&meme.Address)
	if err != nil {
		log.Errorf("%s %s name not known; %s", meme.Type, meme.Address.String(), err.Error())
		return err
	}
	log.Debugf("NFT contract %s name: %s", meme.Address.String(), meme.Name)

	meme.Symbol, err = repo.MemeTokenSymbol(&meme.Address)
	if err != nil {
		log.Errorf("%s %s symbol not known; %s", meme.Type, meme.Address.String(), err.Error())
		return err
	}
	log.Debugf("NFT contract %s symbol: %s", meme.Address.String(), meme.Symbol)

	legacyCollection, err := repo.GetLegacyMemeToken(meme.Address)
	if err != nil {
		log.Errorf("%s %s unable to load off-chain data; %s", meme.Type, meme.Address.String(), err.Error())
		return err
	}

	if nil != legacyCollection {
		meme.Categories, err = legacyCollection.CategoriesAsInt()
		if err != nil {
			log.Errorf("%s %s unable to decode categories; %s", meme.Type, meme.Address.String(), err.Error())
			return err
		}
	}

	/*
		if err = extenMemeTokenMintDetails(meme); err != nil {
			log.Criticalf("failed to extend Meme Token MintDetails %s; %s", meme.Address.String(), err.Error())
			return err
		}
	*/

	return nil
}

/*
func extenMemeTokenMintDetails(nft *types.Collection) (err error) {
	nft.MintDetails = types.CollectionMintDetails{
		PublicMint:    true,        // NB: also need to set fiLegacyCollectionIsOwnerOnly when registering
		IsErc1155:     false,       // unused here..
		HasBaseUri:    false,       // unused here..
		MaxItems:      0,           // TODO: set MaxBlocks
		MaxItemCount:  0,           // unused here..
		MintStartTime: time.Time{}, // actually unused here..
		MintEndTime:   time.Time{}, // actually unused here..
		RevealTime:    time.Time{}, // unused here..
	}

	nft.MemeDetails = types.MemeTokenDetails{
		InitialReserves: big.Int{},
		StakingAmount:   big.Int{},
		BlocksAmount:    big.Int{},
		BlocksFee:       big.Int{},
		BlocksMaxSupply: 0,
	}

	biVal, err := repo.CollectionErc20InitialReserves(&nft.Address)
	if err != nil {
		log.Errorf("%s %s initialReserves not known; %s", nft.Type, nft.Address.String(), err.Error())
	} else {
		nft.MemeDetails.InitialReserves = *biVal
	}
	
	biVal, err := repo.CollectionErc20StakingAmount(&nft.Address)
	if err != nil {
		log.Errorf("%s %s stakingAmount not known; %s", nft.Type, nft.Address.String(), err.Error())
	} else {
		nft.MemeDetails.StakingAmount = *biVal
	}	

	bVal, err := repo.CollectionErc20BlocksAmount(&nft.Address)
	if err != nil {
		log.Errorf("%s %s blocksAmount not known; %s", nft.Type, nft.Address.String(), err.Error())
	} else {
		nft.MemeDetails.BlocksAmount = bVal.Uint64()
	}

	biVal, err = repo.CollectionErc20BlocksFee(&nft.Address)
	if err != nil {
		log.Errorf("%s %s blocksFee not known; %s", nft.Type, nft.Address.String(), err.Error())
	} else {
		nft.MemeDetails.BlocksFee = *biVal
	}

	bVal, err = repo.CollectionErc20BlocksMaxSupply(&nft.Address)
	if err != nil {
		log.Errorf("%s %s blocksMaxSupply not known; %s", nft.Type, nft.Address.String(), err.Error())
	} else {
		nft.MemeDetails.BlocksMaxSupply = bVal.Uint64()
	}

	return nil
}
*/

// newNFTContract handles log event for new factory deployed ERC721/ERC1155 contract.
// Factory::event ContractCreated(address creator, address nft)
// MM Factory::event ContractCreated(address creator, address nft, bool isPrivate)
func newNFTContract(evt *eth.Log, lo *logObserver) {
	// sanity check: no additional topics; 2 x Address = 2 x 32 bytes
	//if len(evt.Data) != 64 || len(evt.Topics) != 1 {
	// sanity check: no additional topics; 3 x Address = 3 x 32 bytes
	if len(evt.Data) != 96 || len(evt.Topics) != 1 {
		log.Errorf("not Factory::ContractCreated() event #%d/#%d; expected 96 bytes of data, %d given; expected 1 topic, %d given",
			evt.BlockNumber, evt.Index, len(evt.Data), len(evt.Topics))
		return
	}

	// make NFT address
	ca := common.Address{}
	ca.SetBytes(evt.Data[32:64])
	nft := types.Collection{
		Address:      ca,
		IsActive:     true,
		IsOurFactory: true,
	}
	log.Infof("found (newNFTContract) NFT contract %s", ca.String())

	// load NFT details
	if err := extendNFTCollectionDetails(&nft, evt, lo); err != nil {
		log.Criticalf("failed to load NFT collection %s; %s", nft.Address.String(), err.Error())
		return
	}

	// add the collection to persistent storage
	if err := repo.StoreCollection(&nft); err != nil {
		log.Criticalf("can not store NFT collection %s; %s", nft.Address.String(), err.Error())
		return
	}

	// add observed contract based on the collection
	addObservedContract(&nft, evt)
	log.Infof("new NFT collection %s found at %s", nft.Name, nft.Address.String())
}

// extendNFTCollectionDetails collects details of an NFT contract.
func extendNFTCollectionDetails(nft *types.Collection, evt *eth.Log, _ *logObserver) (err error) {
	nft.Type, err = repo.NFTContractType(&nft.Address)
	if err != nil {
		log.Errorf("contract %s type not known; %s", evt.Address.String(), err.Error())
		return err
	}
	log.Debugf("NFT contract %s is %s", nft.Address.String(), nft.Type)

	blk, err := repo.GetHeader(evt.BlockNumber)
	if err != nil {
		log.Errorf("header #%d not available; %s", evt.BlockNumber, err.Error())
		return err
	}
	nft.Created = types.Time(time.Unix(int64(blk.Time), 0))

	return extendCollectionMetadata(nft)
}

// extendCollectionMetadata adds metadata to the provided collection structure.
func extendCollectionMetadata(nft *types.Collection) (err error) {
	nft.Name, err = repo.CollectionName(&nft.Address)
	if err != nil {
		log.Errorf("%s %s name not known; %s", nft.Type, nft.Address.String(), err.Error())
		return err
	}
	log.Debugf("NFT contract %s name: %s", nft.Address.String(), nft.Name)

	nft.Symbol, err = repo.CollectionSymbol(&nft.Address)
	if err != nil {
		log.Errorf("%s %s symbol not known; %s", nft.Type, nft.Address.String(), err.Error())
		return err
	}
	log.Debugf("NFT contract %s symbol: %s", nft.Address.String(), nft.Symbol)

	legacyCollection, err := repo.GetLegacyCollection(nft.Address)
	if err != nil {
		log.Errorf("%s %s unable to load off-chain data; %s", nft.Type, nft.Address.String(), err.Error())
		return err
	}

	if nil != legacyCollection {
		nft.Categories, err = legacyCollection.CategoriesAsInt()
		if err != nil {
			log.Errorf("%s %s unable to decode categories; %s", nft.Type, nft.Address.String(), err.Error())
			return err
		}
	}

	/*
		if err = extendNFTCollectionMintDetails(nft); err != nil {
			log.Criticalf("failed to extend NFT Collection MintDetails %s; %s", nft.Address.String(), err.Error())
			return err
		}
	*/

	return nil
}

/*
func extendNFTCollectionMintDetails(nft *types.Collection) (err error) {
	nft.MintDetails = types.CollectionMintDetails{
		PublicMint:    false,                                   // NB: also need to set fiLegacyCollectionIsOwnerOnly when registering
		IsErc1155:     (nft.Type == types.ContractTypeERC1155), // instead of erc721
		HasBaseUri:    false,
		MaxItems:      0,
		MaxItemCount:  0,
		MintStartTime: time.Time{},
		MintEndTime:   time.Time{},
		RevealTime:    time.Time{},
	}

	nft.MemeDetails = types.MemeTokenDetails{
		InitialReserves: big.Int{},
		BlocksAmount:    big.Int{},
		BlocksFee:       big.Int{},
		BlocksMaxSupply: 0,
	}

	if nft.Type == types.ContractTypeERC1155 {
		nft.MintDetails.PublicMint, err = repo.CollectionErc1155IsPrivate(&nft.Address)
		if err != nil {
			log.Errorf("%s %s isPrivate not known; %s", nft.Type, nft.Address.String(), err.Error())
		}
		maxItemCount, err := repo.CollectionErc1155MaxItemSupply(&nft.Address)
		if err != nil {
			log.Errorf("%s %s isPrivate not known; %s", nft.Type, nft.Address.String(), err.Error())
		} else {
			nft.MintDetails.MaxItemCount = maxItemCount.Uint64()
		}
		maxSupply, err := repo.CollectionErc1155MaxSupply(&nft.Address)
		if err != nil {
			log.Errorf("%s %s maxSupply not known; %s", nft.Type, nft.Address.String(), err.Error())
		} else {
			nft.MintDetails.MaxItems = maxSupply.Uint64()
		}
		mTime, err := repo.CollectionErc1155MintStartTime(&nft.Address)
		if err != nil {
			log.Errorf("%s %s mintStartTime not known; %s", nft.Type, nft.Address.String(), err.Error())
		} else {
			nft.MintDetails.MintStartTime = time.Unix(mTime.Int64(), 0)
		}
		mTime, err = repo.CollectionErc1155MintStopTime(&nft.Address)
		if err != nil {
			log.Errorf("%s %s mintEndTime not known; %s", nft.Type, nft.Address.String(), err.Error())
		} else {
			nft.MintDetails.MintEndTime = time.Unix(mTime.Int64(), 0)
		}
		mTime, err = repo.CollectionErc1155RevealTime(&nft.Address)
		if err != nil {
			log.Errorf("%s %s revealTime not known; %s", nft.Type, nft.Address.String(), err.Error())
		} else {
			nft.MintDetails.RevealTime = time.Unix(mTime.Int64(), 0)
		}

	} else {
		nft.MintDetails.PublicMint, err = repo.CollectionErc721IsPrivate(&nft.Address)
		if err != nil {
			log.Errorf("%s %s isPrivate not known; %s", nft.Type, nft.Address.String(), err.Error())
		}
		nft.MintDetails.HasBaseUri, err = repo.CollectionErc721UseBaseUri(&nft.Address)
		if err != nil {
			log.Errorf("%s %s useBaseUri not known; %s", nft.Type, nft.Address.String(), err.Error())
		}
		maxSupply, err := repo.CollectionErc721MaxSupply(&nft.Address)
		if err != nil {
			log.Errorf("%s %s maxSupply not known; %s", nft.Type, nft.Address.String(), err.Error())
		} else {
			nft.MintDetails.MaxItems = maxSupply.Uint64()
		}
		mTime, err := repo.CollectionErc721MintStartTime(&nft.Address)
		if err != nil {
			log.Errorf("%s %s mintStartTime not known; %s", nft.Type, nft.Address.String(), err.Error())
		} else {
			nft.MintDetails.MintStartTime = time.Unix(mTime.Int64(), 0)
		}
		mTime, err = repo.CollectionErc721MintStopTime(&nft.Address)
		if err != nil {
			log.Errorf("%s %s mintEndTime not known; %s", nft.Type, nft.Address.String(), err.Error())
		} else {
			nft.MintDetails.MintEndTime = time.Unix(mTime.Int64(), 0)
		}
		mTime, err = repo.CollectionErc721RevealTime(&nft.Address)
		if err != nil {
			log.Errorf("%s %s revealTime not known; %s", nft.Type, nft.Address.String(), err.Error())
		} else {
			nft.MintDetails.RevealTime = time.Unix(mTime.Int64(), 0)
		}
	}

	return nil
}
*/

// addObservedContract adds new observed contract into repository and log observer.
func addObservedContract(nft *types.Collection, evt *eth.Log) {
	ca := common.Address{}
	if nil != evt.Data && 32 <= len(evt.Data) {
		ca.SetBytes(evt.Data[:32])
	}

	oc := types.ObservedContract{
		Address:     nft.Address,
		Type:        nft.Type,
		Created:     nft.Created,
		Creator:     ca,
		BlockNumber: evt.BlockNumber,
		DeployedBy:  evt.TxHash,
	}

	// store observed contract into the repository
	repo.AddObservedContract(&oc)
}
