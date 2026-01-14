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

// Erc721StartingBlockNumber provides the first important block number for the ERC-721 contract.
// We try to get the first Transfer() event on the contract,
// anything before it is irrelevant for this API.
func (o *Opera) Erc20StartingBlockNumber(adr *common.Address) (uint64, error) {
	// instantiate contract
	erc, err := contracts.NewErc20(*adr, o.ftm)
	if err != nil {
		return 0, err
	}

	// iterate over transfers from zero address (e.g. mint calls)
	iter, err := erc.FilterTransfer(nil, []common.Address{{}}, nil)
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
					iter, err = erc.FilterTransfer(&filterOps, []common.Address{{}}, nil)
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
							iter, err = erc.FilterTransfer(&filterOps, []common.Address{{}}, nil)
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

func (o *Opera) Erc20IsFromFactory(contract *common.Address, block *big.Int) (*common.Address, error) {
	// prepare params
	input, err := o.Erc20Abi().Pack("factory")
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
	res, err := o.Erc20Abi().Unpack("factory", data)
	if err != nil {
		log.Errorf("can not decode contract %s name; %s", contract.String(), err.Error())
		return nil, err
	}
	return abi.ConvertType(res[0], new(common.Address)).(*common.Address), nil
}

func (o *Opera) Erc20InitialReserves(contract *common.Address, block *big.Int) (*big.Int, error) {
	// prepare params
	input, err := o.Erc20Abi().Pack("initialReserves")
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

func (o *Opera) Erc20StakingAmount(contract *common.Address, block *big.Int) (*big.Int, error) {
	// prepare params
	input, err := o.Erc20Abi().Pack("stakingAmount")
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

func (o *Opera) Erc20MintBlocksAmount(contract *common.Address, block *big.Int) (*big.Int, error) {
	// prepare params
	input, err := o.Erc20Abi().Pack("mintBlocksAmount")
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

func (o *Opera) Erc20MintBlocksFee(contract *common.Address, block *big.Int) (*big.Int, error) {
	// prepare params
	input, err := o.Erc20Abi().Pack("mintBlocksFee")
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

func (o *Opera) Erc20MintBlocksSupply(contract *common.Address, block *big.Int) (*big.Int, error) {
	// prepare params
	input, err := o.Erc20Abi().Pack("mintBlocksSupply")
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

func (o *Opera) Erc20MintBlocksMaxSupply(contract *common.Address, block *big.Int) (*big.Int, error) {
	// prepare params
	input, err := o.Erc20Abi().Pack("mintBlocksMaxSupply")
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

func (o *Opera) Erc20MHasSupply(contract *common.Address, block *big.Int) (bool, error) {
	// prepare params
	va, err := o.Erc20MintBlocksSupply(contract, block)
	if err != nil {
		log.Errorf("can not get supply; %s", err.Error())
		return false, err
	}

	ma, err := o.Erc20MintBlocksMaxSupply(contract, block)
	if err != nil {
		log.Errorf("can not get max supply; %s", err.Error())
		return false, err
	}

	return (va.Cmp(ma) < 0), nil
}

// CanMintErc721 checks if the given user can mint a new token on the given NFT contract.
func (o *Opera) Erc20CanMintBlocks(contract *common.Address, user *common.Address, fee *big.Int) (bool, error) {
	// MM: TODO.. adjust to newer Factory contract
	data, err := o.abiVolcano20.Pack("mintBlocks", *user, big.NewInt(1), true)
	if err != nil {
		return false, err
	}

	// use default fee, if not specified
	if fee == nil {
		fee, err = o.Erc20MintBlocksFee(contract, nil)
		if err != nil {
			return false, err
		}
		log.Infof("mint blocks fee for %s is %s", contract.String(), (*hexutil.Big)(fee).String())
	}

	// try to estimate the call
	gas, err := o.ftm.EstimateGas(context.Background(), ethereum.CallMsg{
		From:  *user,
		To:    contract,
		Data:  data,
		Value: fee,
	})
	if err != nil {
		log.Warningf("user %s can not mint on ERC-20 %s; %s", user.String(), contract.String(), err.Error())
		return false, nil
	}

	log.Infof("user %s can mint on ERC-20 %s for %d gas", user.String(), contract.String(), gas)
	return true, nil
}
