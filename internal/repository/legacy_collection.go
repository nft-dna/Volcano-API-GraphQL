package repository

import (
	"artion-api-graphql/internal/types"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

func (p *Proxy) GetLegacyCollection(address common.Address) (*types.LegacyCollection, error) {
	key := "lcol_-" + address.String()
	user, err, _ := p.callGroup.Do(key, func() (interface{}, error) {
		return p.cache.GetLegacyCollection(address, p.shared.GetLegacyCollection)
	})
	return user.(*types.LegacyCollection), err
}

func (p *Proxy) ListLegacyCollections(collectionFilter types.CollectionFilter, cursor types.Cursor, count int, backward bool) (out *types.LegacyCollectionList, err error) {
	if collectionFilter.IsUsed() || cursor != "" || backward {
		return p.shared.ListLegacyCollections(collectionFilter, cursor, count, backward)
	}

	// cache result for the default (mostly used) params
	list, err, _ := p.callGroup.Do("lcols", func() (interface{}, error) {
		return p.cache.GetLegacyCollectionList(func() (*types.LegacyCollectionList, error) {
			return p.shared.ListLegacyCollections(types.CollectionFilter{}, "", 5000, false)
		})
	})

	colList := list.(*types.LegacyCollectionList)
	if colList != nil && len(colList.Collection) > count {
		newColList := *colList // clone callGroup output before modification
		newColList.Collection = newColList.Collection[:count]
		newColList.HasNext = true
		return &newColList, err
	}
	return colList, err
}

func (p *Proxy) UploadCollectionApplication(app types.CollectionApplication, image types.Image, owner common.Address) (err error) {
	// MM
	var imageCid string
	if p.pinner.EmulateOnSharedDb() {
		err = p.shared.StoreImage(&image)
		if err != nil {
			return fmt.Errorf("saving collection image on share_db failed; %s", err)
		}
		imageCid = image.ID().Hex()
	} else {
		imageCid, err = p.pinner.PinFile("collection-image-"+app.Contract.String(), image.Data)
		if err != nil {
			return fmt.Errorf("uploading collection image failed; %s", err)
		}
	}

	mintDetails := types.CollectionMintDetails{
		PublicMint:    false, // NB: also need to set fiLegacyCollectionIsOwnerOnly when registering
		IsErc1155:     false, // instead of erc721
		HasBaseUri:    false,
		MaxItems:      0,
		MaxItemCount:  0,
		MintStartTime: time.Time{},
		MintEndTime:   time.Time{},
		RevealTime:    time.Time{},
	}
	memeDetails := types.MemeTokenDetails{
		InitialReserves: "", //big.Int{},
		StakingAmount:   "", //big.Int{},
		BlocksAmount:    "", //big.Int{},
		BlocksFee:       "", //big.Int{},
		BlocksMaxSupply: 0,
	}

	isOwnerOnly := false
	// check if it is created by 'our factory' contract
	// TODO.. implement an 'interface' or check for contract creator address...
	isInternal := true

	if p.IsErc1155Contract(&app.Contract) {
		isInternal = p.extendErc1155CollectionMintDetails(&app.Contract, &mintDetails)
		if err != nil {
			log.Criticalf("failed to extend Erc1155 Collection MintDetails %s; %s", app.Contract.String(), err.Error())
			return err
		}
		isOwnerOnly = !mintDetails.PublicMint
		mintDetails.IsErc1155 = true
	} else {
		isInternal = p.extendErc721CollectionMintDetails(&app.Contract, &mintDetails)
		if err != nil {
			log.Criticalf("failed to extend Erc721 Collection MintDetails %s; %s", app.Contract.String(), err.Error())
			return err
		}
		isOwnerOnly = !mintDetails.PublicMint
	}
	collection := app.ToCollection(imageCid, &owner, cfg.Server.AddCollectionAsAppropriate, isInternal, !isInternal || isOwnerOnly, mintDetails, memeDetails)

	return p.shared.InsertLegacyCollection(collection)
}

func (p *Proxy) extendErc1155CollectionMintDetails(adr *common.Address, mintDetails *types.CollectionMintDetails) bool {
	isInternal := true
	var err error

	fval, err := p.CollectionErc1155IsFromFactory(adr)
	if err != nil {
		log.Errorf("%s isFromFactory not known; %s", adr.String(), err.Error())
		isInternal = false
	} else {
		log.Infof("%s isFromFactory: %s", adr.String(), fval.String())
	}

	isprivate, err := p.CollectionErc1155IsPrivate(adr)
	if err != nil {
		log.Errorf("%s isPrivate not known; %s", adr.String(), err.Error())
		isInternal = false
	} else {
		mintDetails.PublicMint = !isprivate
	}

	mintDetails.HasBaseUri = true

	maxItemCount, err := p.CollectionErc1155MaxItemSupply(adr)
	if err != nil {
		log.Errorf("%s isPrivate not known; %s", adr.String(), err.Error())
		isInternal = false
	} else {
		mintDetails.MaxItemCount = maxItemCount.Uint64()
	}
	maxSupply, err := p.CollectionErc1155MaxSupply(adr)
	if err != nil {
		log.Errorf("%s maxSupply not known; %s", adr.String(), err.Error())
		isInternal = false
	} else {
		mintDetails.MaxItems = maxSupply.Uint64()
	}
	mTime, err := p.CollectionErc1155MintStartTime(adr)
	if err != nil {
		log.Errorf("%s mintStartTime not known; %s", adr.String(), err.Error())
		isInternal = false
	} else {
		mintDetails.MintStartTime = time.Unix(mTime.Int64(), 0)
	}
	mTime, err = p.CollectionErc1155MintStopTime(adr)
	if err != nil {
		log.Errorf("%s mintEndTime not known; %s", adr.String(), err.Error())
		isInternal = false
	} else {
		mintDetails.MintEndTime = time.Unix(mTime.Int64(), 0)
	}
	mTime, err = p.CollectionErc1155RevealTime(adr)
	if err != nil {
		log.Errorf("%s revealTime not known; %s", adr.String(), err.Error())
		//isInternal = false
	} else {
		mintDetails.RevealTime = time.Unix(mTime.Int64(), 0)
	}
	return isInternal
}

func (p *Proxy) extendErc721CollectionMintDetails(adr *common.Address, mintDetails *types.CollectionMintDetails) bool {
	isInternal := true
	var err error

	fval, err := p.CollectionErc721IsFromFactory(adr)
	if err != nil {
		log.Errorf("%s isFromFactory not known; %s", adr.String(), err.Error())
		//isInternal = false debug .. wip..  previuos versions hadn't this method
	} else {
		log.Infof("%s isFromFactory: %s", adr.String(), fval.String())
	}

	isprivate, err := p.CollectionErc721IsPrivate(adr)
	if err != nil {
		log.Errorf("%s isPrivate not known; %s", adr.String(), err.Error())
		isInternal = false
	} else {
		mintDetails.PublicMint = !isprivate
	}
	mintDetails.HasBaseUri, err = p.CollectionErc721UseBaseUri(adr)
	if err != nil {
		log.Errorf("%s useBaseUri not known; %s", adr.String(), err.Error())
		//isInternal = false
	}
	maxSupply, err := p.CollectionErc721MaxSupply(adr)
	if err != nil {
		log.Errorf("%s maxSupply not known; %s", adr.String(), err.Error())
		isInternal = false
	} else {
		mintDetails.MaxItems = maxSupply.Uint64()
	}
	mTime, err := p.CollectionErc721MintStartTime(adr)
	if err != nil {
		log.Errorf("%s mintStartTime not known; %s", adr.String(), err.Error())
		isInternal = false
	} else {
		mintDetails.MintStartTime = time.Unix(mTime.Int64(), 0)
	}
	mTime, err = p.CollectionErc721MintStopTime(adr)
	if err != nil {
		log.Errorf("%s mintEndTime not known; %s", adr.String(), err.Error())
		isInternal = false
	} else {
		mintDetails.MintEndTime = time.Unix(mTime.Int64(), 0)
	}
	mTime, err = p.CollectionErc721RevealTime(adr)
	if err != nil {
		log.Errorf("%s revealTime not known; %s", adr.String(), err.Error())
		//isInternal = false
	} else {
		mintDetails.RevealTime = time.Unix(mTime.Int64(), 0)
	}
	return isInternal
}

// MustCollectionName provides a name of an Artion ERC721 and/or ERC1155 token,
// or an empty string, if the name is not available.
func (p *Proxy) MustCollectionName(adr *common.Address) string {
	c, err := p.shared.GetLegacyCollection(*adr)
	if err != nil || c == nil {
		return adr.String()
	}
	if c.Name == "" {
		return adr.String()
	}
	return c.Name
}

func (p *Proxy) ApproveCollection(address common.Address) (err error) {
	err = p.shared.ApproveCollection(address)
	p.cache.InvalidateLegacyCollection(address)
	return err
}

func (p *Proxy) DeclineCollection(address common.Address) (err error) {
	err = p.shared.DeclineCollection(address)
	p.cache.InvalidateLegacyCollection(address)
	return err
}

func (p *Proxy) BanCollection(address common.Address) (err error) {
	err = p.shared.BanCollection(address)
	p.cache.InvalidateLegacyCollection(address)
	return err
}

func (p *Proxy) UnbanCollection(address common.Address) (err error) {
	err = p.shared.UnbanCollection(address)
	p.cache.InvalidateLegacyCollection(address)
	return err
}

func (p *Proxy) ListCollectionsWithAppropriateUpdate(after time.Time, maxAmount int64) (out []*types.LegacyCollection, err error) {
	return p.shared.ListCollectionsWithAppropriateUpdate(after, maxAmount)
}

// IsCollectionAppropriate checks if given collection is approved/not banned collection.
func (p *Proxy) IsCollectionAppropriate(contract *common.Address) bool {
	return p.shared.IsCollectionAppropriate(contract)
}
