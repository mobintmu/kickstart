package campaign

import (
	"context"
	"fmt"
	"kickstart/contracts"
	"kickstart/service/account"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Campaign struct {
	instance  *contracts.Campaign
	txAddress common.Address
	tx        *types.Transaction
	Min       *big.Int
	client    *ethclient.Client
	chainID   *big.Int
}

func NewCampaign(min *big.Int, client *ethclient.Client) *Campaign {

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	return &Campaign{
		Min:     min,
		client:  client,
		chainID: chainID,
	}
}

func (c *Campaign) Deploy(manager *account.Account) {
	nonce := manager.GetNonce()

	gasPrice, err := c.client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// chainID, err := c.client.ChainID(context.Background())
	// if err != nil {
	// 	log.Fatal(err)
	// }

	auth, err := bind.NewKeyedTransactorWithChainID(manager.PrivateKey, c.chainID)
	if err != nil {
		fmt.Println(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	txAddress, tx, instance, err := contracts.DeployCampaign(auth, c.client, c.Min)
	if err != nil {
		fmt.Println("DeployLottery")
		log.Fatal(err.Error())
	}

	c.instance = instance
	c.txAddress = txAddress
	c.tx = tx

	fmt.Println(txAddress.Hex())
	fmt.Println(tx.Hash().Hex())

}

func (c *Campaign) Contribute(contributor *account.Account) {

	nonce := contributor.GetNonce()

	gasPrice, err := c.client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// chainID, err := c.client.ChainID(context.Background())
	// if err != nil {
	// 	log.Fatal(err)
	// }

	auth, err := bind.NewKeyedTransactorWithChainID(contributor.PrivateKey, c.chainID)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = c.Min              // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice

	tx, err := c.instance.Contribute(auth)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("transaction hash ...")
	fmt.Println(tx.Hash())

}

func (c *Campaign) IsPaid(contributor *account.Account) bool {

	result, err := c.instance.CampaignCaller.IsPaid(&bind.CallOpts{
		Pending:     false,
		From:        contributor.PublicAddress,
		BlockNumber: &big.Int{},
		BlockHash:   [32]byte{},
		Context:     nil,
	})
	if err != nil {
		fmt.Println("IsPaid")
		log.Fatal(err)
	}

	result, err = c.instance.IsPaid(&bind.CallOpts{
		Pending:     false,
		From:        contributor.PublicAddress,
		BlockNumber: &big.Int{},
		BlockHash:   [32]byte{},
		Context:     nil,
	})

	fmt.Println(result)

	return result

}

func (c *Campaign) GetContributor(contributor *account.Account) bool {

	result, err := c.instance.Approvers(nil, contributor.PublicAddress)
	if err != nil {
		fmt.Println("GetContributor")
		log.Fatal(err)
	}
	fmt.Println(result)

	result, err = c.instance.Approvers(&bind.CallOpts{
		Pending:     false,
		From:        contributor.PublicAddress,
		BlockNumber: &big.Int{},
		BlockHash:   [32]byte{},
		Context:     nil,
	}, contributor.PublicAddress)
	if err != nil {
		fmt.Println("GetContributor")
		log.Fatal(err)
	}
	fmt.Println(result)

	return result
}
