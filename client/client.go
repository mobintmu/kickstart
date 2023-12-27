package client

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

type Client struct {
	Client *ethclient.Client
}

func NewClient(address string) *Client {

	client, err := ethclient.Dial(address)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connect to network ...")

	newClient := &Client{
		Client: client,
	}


	return newClient
}
