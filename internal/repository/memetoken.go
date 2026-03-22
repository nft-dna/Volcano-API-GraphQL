// Package repository implements persistent data access and processing.
package repository

import (
	"artion-api-graphql/internal/types"
	"bytes"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// addCollectionQueueCapacity is the capacity of the queue being filled with
// new collections to be added into the repository.
const addMemeTokenQueueCapacity = 10

// NewCollectionQueue provides queue filled with addresses of collection contracts
// to be added into the API.
func (p *Proxy) NewMemeTokenQueue() chan common.Address {
	return p.newMemeTokenQueue
}

// AddNewCollection pushes the given address to be validated and added as a new collection.
func (p *Proxy) AddNewMemeToken(adr *common.Address) {
	p.newMemeTokenQueue <- *adr
}

// CollectionName provides a name of an Artion ERC721 and/or ERC1155 token.
func (p *Proxy) MemeTokenName(adr *common.Address) (string, error) {
	return p.rpc.MemeTokenName(adr)
}

// CollectionSymbol provides a symbol of an Artion ERC721 and/or ERC1155 token.
func (p *Proxy) MemeTokenSymbol(adr *common.Address) (string, error) {
	return p.rpc.MemeTokenSymbol(adr)
}

// StoreCollection adds the specified NFT collection contract record into persistent storage.
func (p *Proxy) StoreMemeToken(nft *types.Collection) error {
	return p.db.StoreMemeToken(nft)
}

func (p *Proxy) InsertLegacyMemeToken(c types.LegacyCollection, isUpload bool) error {
	return p.shared.InsertLegacyMemeToken(c, isUpload)
}

func (p *Proxy) GetMemeToken(address common.Address) (*types.Collection, error) {
	key := "Col-" + address.String()
	user, err, _ := p.callGroup.Do(key, func() (interface{}, error) {
		return p.db.GetMemeToken(address)
	})
	return user.(*types.Collection), err
}

func (p *Proxy) ListMemeTokens(cursor types.Cursor, count int, backward bool) (*types.CollectionList, error) {
	return p.db.ListMemeTokens(cursor, count, backward)
}

func (p *Proxy) MemeBlocksSupply(contract *common.Address) (*big.Int, error) {
	return p.rpc.MemeBlocksSupply(contract)
}

func (p *Proxy) IncMemeBlocksSupply(contract *common.Address, count uint64) bool {
	return p.shared.IncMemeBlocksSupply(*contract, count)
}

// CanMint checks if the given user can mint a new token on the given NFT contract.
func (p *Proxy) CanMintBlock(contract *common.Address, user *common.Address, fee *big.Int) (bool, error) {
	return p.rpc.CanMintMemeBlocks(contract, user, fee)
}

// CollectionOwner tries to get the owner of the given collection.
func (p *Proxy) MemeTokenOwner(contract *common.Address) *common.Address {
	return p.rpc.MemeTokenOwner(contract)
}

// CanRegisterCollection checks if the given collection can be registered.
func (p *Proxy) CanRegisterMemeToken(contract *common.Address, user *common.Address) error {

	if !p.IsErc20Contract(contract) {
		return fmt.Errorf("the contract is not ERC-20")
	}

	owner := p.rpc.MemeTokenOwner(contract)
	if owner != nil {
		if !bytes.Equal(user.Bytes(), owner.Bytes()) {
			return fmt.Errorf("the contract is owned by somebody else")
		}
	}

	return nil
}
