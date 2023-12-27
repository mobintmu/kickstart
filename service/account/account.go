package account

import (
	"context"
	"kickstart/client"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Account struct {
	PublicAddress  common.Address
	PrivateAddress string
	Name           string
	client         *ethclient.Client
}

// NewAccount creates a new instance of Account
func NewAccount(publicAddress string,
	privateAddress string,
	client *client.Client) *Account {
	return &Account{
		PublicAddress: common.HexToAddress(publicAddress),
		client:        client.Client,
	}
}

func (a *Account) GetBalance() *big.Int {

	balance, err := a.client.BalanceAt(context.Background(), a.PublicAddress, nil)
	if err != nil {
		log.Fatal(err)
	}

	return balance

}
