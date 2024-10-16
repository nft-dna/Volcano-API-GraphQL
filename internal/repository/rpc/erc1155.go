// Package rpc provides high level access to the Volcano Opera blockchain
// node through RPC interface.
package rpc

import (
	"artion-api-graphql/internal/repository/rpc/contracts"
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Erc1155BalanceOf extracts balance of a NFT for an owner.
func (o *Opera) Erc1155BalanceOf(contract *common.Address, tokenId *big.Int, owner *common.Address, block *big.Int) (*big.Int, error) {
	// prepare params
	input, err := o.Erc1155Abi().Pack("balanceOf", *owner, tokenId)
	if err != nil {
		log.Errorf("can not pack data; %s", err.Error())
		return nil, err
	}

	// call the contract
	data, err := o.ftm.CallContract(context.Background(), ethereum.CallMsg{
		From: common.Address{},
		To:   contract,
		Data: input,
	}, block)
	if err != nil {
		return nil, err
	}
	return new(big.Int).SetBytes(data), nil
}

func (o *Opera) Erc1155IsFromFactory(contract *common.Address, block *big.Int) (*common.Address, error) {
	// prepare params
	input, err := o.Erc1155Abi().Pack("factory")
	if err != nil {
		log.Errorf("can not pack data; %s", err.Error())
		return nil, err
	}

	// call the contract
	data, err := o.ftm.CallContract(context.Background(), ethereum.CallMsg{
		From: common.Address{},
		To:   contract,
		Data: input,
	}, block)
	if err != nil {
		return nil, err
	}
	res, err := o.Erc1155Abi().Unpack("factory", data)
	if err != nil {
		log.Errorf("can not decode contract %s name; %s", contract.String(), err.Error())
		return nil, err
	}
	return abi.ConvertType(res[0], new(common.Address)).(*common.Address), nil
}

func (o *Opera) Erc1155IsPrivate(contract *common.Address, block *big.Int) (bool, error) {
	// prepare params
	input, err := o.Erc1155Abi().Pack("isprivate")
	if err != nil {
		log.Errorf("can not pack data; %s", err.Error())
		return false, err
	}

	// call the contract
	data, err := o.ftm.CallContract(context.Background(), ethereum.CallMsg{
		From: common.Address{},
		To:   contract,
		Data: input,
	}, block)
	if err != nil {
		return false, err
	}
	return len(data) == 32 && data[0] == 0 && data[31] > 0, nil
}

func (o *Opera) Erc1155TotalSupply(contract *common.Address, block *big.Int) (*big.Int, error) {
	// prepare params
	input, err := o.Erc1155Abi().Pack("itemsSupply")
	if err != nil {
		log.Errorf("can not pack data; %s", err.Error())
		return nil, err
	}

	// call the contract
	data, err := o.ftm.CallContract(context.Background(), ethereum.CallMsg{
		From: common.Address{},
		To:   contract,
		Data: input,
	}, block)
	if err != nil {
		return nil, err
	}
	return new(big.Int).SetBytes(data), nil
}

func (o *Opera) Erc1155MaxSupply(contract *common.Address, block *big.Int) (*big.Int, error) {
	// prepare params
	input, err := o.Erc1155Abi().Pack("maxSupply")
	if err != nil {
		log.Errorf("can not pack data; %s", err.Error())
		return nil, err
	}

	// call the contract
	data, err := o.ftm.CallContract(context.Background(), ethereum.CallMsg{
		From: common.Address{},
		To:   contract,
		Data: input,
	}, block)
	if err != nil {
		return nil, err
	}
	return new(big.Int).SetBytes(data), nil
}

func (o *Opera) Erc1155MaxItemSupply(contract *common.Address, block *big.Int) (*big.Int, error) {
	// prepare params
	input, err := o.Erc1155Abi().Pack("maxItemSupply")
	if err != nil {
		log.Errorf("can not pack data; %s", err.Error())
		return nil, err
	}

	// call the contract
	data, err := o.ftm.CallContract(context.Background(), ethereum.CallMsg{
		From: common.Address{},
		To:   contract,
		Data: input,
	}, block)
	if err != nil {
		return nil, err
	}
	return new(big.Int).SetBytes(data), nil
}

func (o *Opera) Erc1155MintStartTime(contract *common.Address, block *big.Int) (*big.Int, error) {
	// prepare params
	input, err := o.Erc1155Abi().Pack("mintStartTime")
	if err != nil {
		log.Errorf("can not pack data; %s", err.Error())
		return nil, err
	}

	// call the contract
	data, err := o.ftm.CallContract(context.Background(), ethereum.CallMsg{
		From: common.Address{},
		To:   contract,
		Data: input,
	}, block)
	if err != nil {
		return nil, err
	}
	return new(big.Int).SetBytes(data), nil
}

func (o *Opera) Erc1155MintStopTime(contract *common.Address, block *big.Int) (*big.Int, error) {
	// prepare params
	input, err := o.Erc1155Abi().Pack("mintStopTime")
	if err != nil {
		log.Errorf("can not pack data; %s", err.Error())
		return nil, err
	}

	// call the contract
	data, err := o.ftm.CallContract(context.Background(), ethereum.CallMsg{
		From: common.Address{},
		To:   contract,
		Data: input,
	}, block)
	if err != nil {
		return nil, err
	}
	return new(big.Int).SetBytes(data), nil
}

func (o *Opera) Erc1155RevealTime(contract *common.Address, block *big.Int) (*big.Int, error) {
	// prepare params
	input, err := o.Erc1155Abi().Pack("revealTime")
	if err != nil {
		log.Errorf("can not pack data; %s", err.Error())
		return nil, err
	}

	// call the contract
	data, err := o.ftm.CallContract(context.Background(), ethereum.CallMsg{
		From: common.Address{},
		To:   contract,
		Data: input,
	}, block)
	if err != nil {
		return nil, err
	}
	return new(big.Int).SetBytes(data), nil
}

// Erc1155TokenUri gets a token specific URI address from ERC-1155 contract using uri() call.
func (o *Opera) Erc1155TokenUri(contract *common.Address, tokenId *big.Int) (string, error) {
	// prepare params
	input, err := o.Erc1155Abi().Pack("uri", tokenId)
	if err != nil {
		log.Errorf("can not pack data; %s", err.Error())
		return "", err
	}

	// call the contract
	data, err := o.ftm.CallContract(context.Background(), ethereum.CallMsg{
		From: common.Address{},
		To:   contract,
		Data: input,
	}, nil)
	res, err := o.abiVolcano721.Unpack("uri", data)
	if err != nil {
		log.Errorf("can not decode response; %s", err.Error())
		return "", err
	}

	// extract the URI from the response data; make sure to follow ERC-1155 specs for token ID replacement
	// If the string {id} exists in the URI, client MUST replace this
	// with the actual token ID in hexadecimal form in lower case, not prefixed, zero-padded to 64 hex chars.
	// https://eips.ethereum.org/EIPS/eip-1155#metadata
	return strings.Replace(
		*abi.ConvertType(res[0], new(string)).(*string),
		"{id}",
		fmt.Sprintf("%064x", tokenId), -1), nil
}

// Erc1155StartingBlockNumber provides the first important block number for the ERC-1155 contract.
func (o *Opera) Erc1155StartingBlockNumber(adr *common.Address) (uint64, error) {
	// instantiate contract
	erc, err := contracts.NewErc1155(*adr, o.ftm)
	if err != nil {
		return 0, err
	}

	bSingle, err := o.Erc1155FirstMintBlock(erc)
	if err != nil {
		return 0, err
	}

	bBatch, err := o.Erc1155FirstMintBatchBlock(erc)
	if err != nil {
		return 0, err
	}

	if bSingle > bBatch {
		return bSingle, nil
	}
	return bBatch, nil
}

// Erc1155FirstMintBlock tries to find the first block the ERC-1155 contract
// was used to mint a single token.
func (o *Opera) Erc1155FirstMintBlock(erc *contracts.Erc1155) (uint64, error) {
	iter, err := erc.FilterTransferSingle(nil, nil, []common.Address{{}}, nil)
	if err != nil {

		// MM timeout on Sepolia 'open' rpc
		if strings.Contains(strings.ToLower(fmt.Sprint(err)), "timed out") {
			// do something
			chb, perr := o.CurrentHead()
			if perr == nil {
				step := uint64(20000)
				resOk := false
				b := o.minBlockNumber
				for b <= chb {
					stop := b + step
					filterOps := bind.FilterOpts{
						Context: context.Background(),
						Start:   b,
						End:     &stop,
					}
					iter, err = erc.FilterTransferSingle(&filterOps, []common.Address{{}}, nil, nil)
					if err == nil && iter.Event != nil {
						resOk = true
						break
					}
					b = b + step
				}
				/*
					if !resOk {
						b = o.minBlockNumber
						for b > 0 {

							filterOps := bind.FilterOpts{
								Context: context.Background(),
								Start:   b - step,
								End:     &b,
							}
							iter, err = erc.FilterTransferSingle(&filterOps, []common.Address{{}}, nil, nil)
							if err == nil && iter.Event != nil {
								resOk = true
								break
							}
							if b > step {
								b = b - step
							} else if b > 0 {
								b = 0
							} else {
								break
							}
						}
					}
				*/
				if !resOk {
					return 0, err
				}
			}
		}
	}

	var blk uint64
	if iter.Next() {
		blk = iter.Event.Raw.BlockNumber
	}

	if err := iter.Close(); err != nil {
		log.Errorf("could not close filter iterator; %s", err.Error())
	}
	return blk, nil
}

// Erc1155FirstMintBatchBlock tries to find the first block the ERC-1155 contract
// was used to mint a batch of tokens.
func (o *Opera) Erc1155FirstMintBatchBlock(erc *contracts.Erc1155) (uint64, error) {
	iter, err := erc.FilterTransferBatch(nil, nil, []common.Address{{}}, nil)
	if err != nil {

		// MM timeout on Sepolia 'open' rpc
		if strings.Contains(strings.ToLower(fmt.Sprint(err)), "timed out") {
			// do something
			chb, perr := o.CurrentHead()
			if perr == nil {
				step := uint64(20000)
				resOk := false
				b := o.minBlockNumber
				for b <= chb {
					stop := b + step
					filterOps := bind.FilterOpts{
						Context: context.Background(),
						Start:   b,
						End:     &stop,
					}
					iter, err = erc.FilterTransferBatch(&filterOps, []common.Address{{}}, nil, nil)
					if err == nil && iter.Event != nil {
						resOk = true
						break
					}
					b = b + step
				}
				/*
					if !resOk {
						b = o.minBlockNumber
						for b > 0 {

							filterOps := bind.FilterOpts{
								Context: context.Background(),
								Start:   b - step,
								End:     &b,
							}
							iter, err = erc.FilterTransferBatch(&filterOps, []common.Address{{}}, nil, nil)
							if err == nil && iter.Event != nil {
								resOk = true
								break
							}
							if b > step {
								b = b - step
							} else if b > 0 {
								b = 0
							} else {
								break
							}
						}
					}
				*/
				if !resOk {
					return 0, err
				}
			}
		}
	}

	var blk uint64
	for iter.Next() {
		if blk < iter.Event.Raw.BlockNumber {
			blk = iter.Event.Raw.BlockNumber
		}
	}

	if err := iter.Close(); err != nil {
		log.Errorf("could not close filter iterator; %s", err.Error())
	}
	return blk, nil
}

// CanMintErc1155 checks if the given user can mint a new token on the given NFT contract.
func (o *Opera) CanMintErc1155(contract *common.Address, user *common.Address, fee *big.Int) (bool, error) {
	// MM: TODO.. adjust to newer Factory contract
	data, err := o.abiVolcano1155.Pack("mint", *user, big.NewInt(1), []byte(defaultMintingTestTokenUrl))
	if err != nil {
		return false, nil
	}

	// use default fee, if not specified
	if fee == nil {
		fee = o.MustPlatformFee(contract)
		log.Infof("platform fee plus creator fee for %s is %s", contract.String(), (*hexutil.Big)(fee).String())
	}

	// try to estimate the call
	gas, err := o.ftm.EstimateGas(context.Background(), ethereum.CallMsg{
		From:  *user,
		To:    contract,
		Data:  data,
		Value: fee,
	})
	if err != nil {
		log.Warningf("user %s can not mint on ERC-1155 %s; %s", user.String(), contract.String(), err.Error())
		return false, nil
	}

	log.Infof("user %s can mint on ERC-1155 %s for %d gas", user.String(), contract.String(), gas)
	return true, nil
}
