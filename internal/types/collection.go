// Package types provides high level structures for the API server.
package types

import (
	"github.com/ethereum/go-ethereum/common"
)

/*
type CollectionMintDetails struct {
	PublicMint    bool      `bson:"publicMint"`
	IsErc1155     bool      `bson:"isErc1155"`
	HasBaseUri    bool      `bson:"hasBaseUri"`
	MaxItems      uint64    `bson:"maxItems"`
	MaxItemCount  uint64    `bson:"maxItemCount"`
	MintStartTime time.Time `bson:"mintStartTime"`
	MintEndTime   time.Time `bson:"mintEndTime"`
	RevealTime    time.Time `bson:"revealTime"`
}

type MemeTokenDetails struct {
	InitialReserves big.Int `bson:"initialReserves"`
	StakingAmount   big.Int `bson:"stakingAmount"`
	BlocksAmount    big.Int  `bson:"blocksAmount"`
	BlocksFee       big.Int `bson:"blocksFee"`
	BlocksMaxSupply uint64  `bson:"blocksMaxSupply"`
}
*/

// Collection represents an Artion token collection, represented by an NFT contract.
// Artion basically recognizes NFT contracts deployed form a designated factory.
type Collection struct {
	Address    common.Address  `bson:"_id"`
	Type       string          `bson:"type"`
	Name       string          `bson:"name"`
	Symbol     string          `bson:"symbol"`
	Created    Time            `bson:"created"`
	Categories []int32         `bson:"categories"`
	IsActive   bool            `bson:"is_active"`
	VerifiedBy *common.Address `bson:"verified_by"`
	//
	IsOurFactory bool `bson:"isOurFactory"`
	//MintDetails  CollectionMintDetails `bson:"mintDetails"`
	//MemeDetails  MemeTokenDetails      `bson:"memeDetails"`
}
