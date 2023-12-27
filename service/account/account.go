package account

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"kickstart/client"
	"kickstart/contracts"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Account struct {
	PublicAddress common.Address
	privateKey    *ecdsa.PrivateKey
	Name          string
	client        *ethclient.Client
}

// NewAccount creates a new instance of Account
func NewAccount(publicAddress string,
	privateAddress string,
	client *client.Client) *Account {

	privateKey, err := crypto.HexToECDSA(strings.TrimPrefix(privateAddress, "0x"))
	if err != nil {
		log.Fatal(err)
	}

	return &Account{
		PublicAddress: common.HexToAddress(publicAddress),
		client:        client.Client,
		privateKey:    privateKey,
	}
}

func (a *Account) GetBalance() *big.Int {

	balance, err := a.client.BalanceAt(context.Background(), a.PublicAddress, nil)
	if err != nil {
		log.Fatal(err)
		return &big.Int{}
	}

	return balance

}

func (a *Account) GetNonce() uint64 {

	nonce, err := a.client.PendingNonceAt(context.Background(), a.PublicAddress)
	if err != nil {
		log.Fatal(err)
		return 0
	}

	return nonce

}

func (a *Account) Deploy(min *big.Int) (common.Address, *types.Transaction) {

	nonce := a.GetNonce()

	gasPrice, err := a.client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	chainID, err := a.client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(a.privateKey, chainID)
	if err != nil {
		fmt.Println(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	txAddress, tx, _, err := contracts.DeployCampaign(auth, a.client, min)
	if err != nil {
		fmt.Println("DeployLottery")
		log.Fatal(err.Error())
	}

	fmt.Println(txAddress.Hex())
	fmt.Println(tx.Hash().Hex())
	return txAddress, tx

}
