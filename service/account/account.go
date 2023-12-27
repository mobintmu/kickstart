package account

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"kickstart/client"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Account struct {
	PublicAddress common.Address
	PrivateKey    *ecdsa.PrivateKey
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
		PrivateKey:    privateKey,
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
func (a *Account) GetHeader() *big.Int {

	header, err := a.client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(header.Number.String()) // 5671744

	return header.Number
}

func (a *Account) GetDataFromBlock(blockNumber *big.Int) {
	block, err := a.client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("block.Number().Uint64() ", block.Number().Uint64())        // 5671744
	fmt.Println("block.Time()", block.Time())                               // 1527211625
	fmt.Println("block.Difficulty().Uint64()", block.Difficulty().Uint64()) // 3217000136609065
	fmt.Println("block.Hash().Hex()", block.Hash().Hex())                   // 0x9e8751ebb5069389b855bba72d94902cc385042661498a415979b7b6ee9ba4b9
	fmt.Println("len(block.Transactions())", len(block.Transactions()))     // 144
}
