package client

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

type Client struct {
	Client *ethclient.Client
}

func (c *Client) NewClient(address string) {

	client, err := ethclient.Dial(address)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connect to network ...")
	c.Client = client
}
