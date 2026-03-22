// Package rpc provides high level access to the Volcano Opera blockchain
// node through RPC interface.
package rpc

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

// MemeTokenName provides a name of an Artion ERC721 and/or ERC1155 token.
// Solidity: function name() view returns(string)
func (o *Opera) MemeTokenName(adr *common.Address) (string, error) {
	data, err := o.ftm.CallContract(context.Background(), ethereum.CallMsg{
		From: common.Address{},
		To:   adr,
		Data: common.Hex2Bytes("06fdde03"),
	}, nil)
	if err != nil {
		log.Errorf("contract %s name not found", adr.String())
		return "", err
	}
	res, err := o.abiVolcano721.Unpack("name", data)
	if err != nil {
		log.Errorf("can not decode contract %s name; %s", adr.String(), err.Error())
		return "", err
	}
	return *abi.ConvertType(res[0], new(string)).(*string), nil
}

// CollectionSymbol provides a symbol of an Artion ERC721 and/or ERC1155 token.
// Solidity: function symbol() view returns(string)
func (o *Opera) MemeTokenSymbol(adr *common.Address) (string, error) {
	data, err := o.ftm.CallContract(context.Background(), ethereum.CallMsg{
		From: common.Address{},
		To:   adr,
		Data: common.Hex2Bytes("95d89b41"),
	}, nil)
	if err != nil {
		log.Errorf("contract %s symbol not found", adr.String())
		return "", err
	}
	res, err := o.abiVolcano721.Unpack("symbol", data)
	if err != nil {
		log.Errorf("can not decode contract %s symbol; %s", adr.String(), err.Error())
		return "", err
	}
	return *abi.ConvertType(res[0], new(string)).(*string), nil
}

// CollectionOwner tries to get the owner of the given collection.
// The call uses public owner access, if available and returns nil if the owner is not known.
// Solidity: function owner() view returns(address)
func (o *Opera) MemeTokenOwner(contract *common.Address) *common.Address {
	data, err := o.ftm.CallContract(context.Background(), ethereum.CallMsg{
		From: common.Address{},
		To:   contract,
		Data: common.Hex2Bytes("8da5cb5b"),
	}, nil)
	if err != nil {
		log.Warningf("owner of %s not available; %s", contract.String(), err.Error())
		return nil
	}

	// we expect 32 bytes of address in return
	if len(data) != 32 {
		log.Warningf("owner of %s not loaded; %d bytes given, expected 32 bytes", contract.String(), len(data))
		return nil
	}

	owner := common.BytesToAddress(data)
	return &owner
}

func (o *Opera) CanMintMemeBlocks(contract *common.Address, user *common.Address, fee *big.Int) (bool, error) {
	return o.Erc20CanMintBlocks(contract, user, fee)
}

func (o *Opera) MemeBlocksSupply(contract *common.Address) (*big.Int, error) {
	return o.Erc20MintBlocksSupply(contract, nil)
}
