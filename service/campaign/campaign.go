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
}

func NewCampaign(min *big.Int, client *ethclient.Client) *Campaign {
	return &Campaign{
		Min:    min,
		client: client,
	}
}

func (c *Campaign) Deploy(manager *account.Account) {
	nonce := manager.GetNonce()

	gasPrice, err := c.client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	chainID, err := c.client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(manager.PrivateKey, chainID)
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

	chainID, err := c.client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(contributor.PrivateKey, chainID)
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
