package main

import (
	"fmt"
	"kickstart/client"
	"kickstart/helper"
	"kickstart/service/account"
	"math/big"

	"github.com/joho/godotenv"
)

func main() {

	// Load the .env file in the current directory
	godotenv.Load()

	address := helper.GetEnvVariable("ADDRESS")

	client := client.NewClient(address)
	defer client.Client.Close()

	manager := account.NewAccount(
		helper.GetEnvVariable("MANAGER_ADDRESS"),
		helper.GetEnvVariable("MANAGER_PRIVATE_ADDRESS"),
		client)
	fmt.Println("balance manager :", manager.GetBalance())

	provider := account.NewAccount(
		helper.GetEnvVariable("PROVIDER_ADDRESS"),
		helper.GetEnvVariable("PROVIDER_PRIVATE_ADDRESS"),
		client)
	fmt.Println("balance provider :", provider.GetBalance())

	contributor := account.NewAccount(
		helper.GetEnvVariable("CONTRIBUTOR_ADDRESS"),
		helper.GetEnvVariable("CONTRIBUTOR_PRIVATE_ADDRESS"),
		client)
	fmt.Println("contributor manager :", contributor.GetBalance())

	fmt.Println("nonce manager : ", manager.GetNonce())

	fmt.Println("deploying contract  ... ")
	manager.Deploy(big.NewInt(300))
	fmt.Println("balance manager :", manager.GetBalance())

}
