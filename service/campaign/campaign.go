package campaign

import (
	"context"
	"fmt"
	"kickstart/contracts"
	"kickstart/entity"
	"kickstart/service/account"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Campaign struct {
	Instance  *contracts.Campaign
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

	chainID, err := c.client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(manager.PrivateKey, chainID)
	if err != nil {
		fmt.Println(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(6721975) // in units
	auth.GasPrice = gasPrice

	txAddress, tx, instance, err := contracts.DeployCampaign(auth, c.client, c.Min)
	if err != nil {
		fmt.Println("DeployLottery")
		fmt.Println(err.Error())
		log.Fatal("stop")
	}

	c.Instance = instance
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

	auth, err := bind.NewKeyedTransactorWithChainID(contributor.PrivateKey, c.chainID)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = c.Min              // in wei
	auth.GasLimit = uint64(6721975) // in units 6721975
	auth.GasPrice = gasPrice

	tx, err := c.Instance.Contribute(auth)
	if err != nil {
		fmt.Println("Contribute")
		fmt.Println(err.Error())
		log.Fatal(err)
	}

	fmt.Println("transaction hash ...")
	fmt.Println(tx.Hash())

}

func (c *Campaign) GetBalance() *big.Int {
	result, err := c.Instance.GetBalance(nil)
	if err != nil {
		fmt.Println("IsPaid")
		fmt.Println(err)
	}

	fmt.Println(result)

	return result
}

func (c *Campaign) IsPaid(contributor *account.Account) bool {

	result, err := c.Instance.IsPaid(&bind.CallOpts{
		From:    contributor.PublicAddress,
		Context: context.Background(),
	})
	if err != nil {
		fmt.Println("IsPaid")
		fmt.Println(err)
	}

	fmt.Println(result)

	return result
}

func (c *Campaign) GetContributor(contributor *account.Account) bool {

	result, err := c.Instance.Approvers(nil, contributor.PublicAddress)
	if err != nil {
		fmt.Println("GetContributor")
		log.Fatal(err)
	}
	fmt.Println(result)

	result, err = c.Instance.Approvers(&bind.CallOpts{
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

func (c *Campaign) AddRequest(request entity.Request, manager *account.Account) {
	nonce := manager.GetNonce()

	gasPrice, err := c.client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(manager.PrivateKey, c.chainID)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(6721975) // in units 6721975
	auth.GasPrice = gasPrice

	tx, err := c.Instance.AddRequest(auth, request.Description, request.Value, request.ProviderAddress)
	if err != nil {
		fmt.Println("CreateRequest")
		fmt.Println(err.Error())
		log.Fatal(err)
	}

	fmt.Println("transaction hash ...")
	fmt.Println(tx.Hash())

}
