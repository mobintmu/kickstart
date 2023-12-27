package main

import (
	"fmt"
	"kickstart/client"
	"kickstart/entity"
	"kickstart/helper"
	"kickstart/service/account"
	"kickstart/service/campaign"
	"math/big"

	"github.com/joho/godotenv"
)

func main() {

	// Load the .env file in the current directory
	godotenv.Load()

	address := helper.GetEnvVariable("ADDRESS")

	client := client.NewClient(address)
	defer client.Client.Close()

	campaign := campaign.NewCampaign(big.NewInt(1000), client.Client)

	manager := account.NewAccount(
		helper.GetEnvVariable("MANAGER_ADDRESS"),
		helper.GetEnvVariable("MANAGER_PRIVATE_ADDRESS"),
		client)

	provider := account.NewAccount(
		helper.GetEnvVariable("PROVIDER_ADDRESS"),
		helper.GetEnvVariable("PROVIDER_PRIVATE_ADDRESS"),
		client)

	contributor := account.NewAccount(
		helper.GetEnvVariable("CONTRIBUTOR_ADDRESS"),
		helper.GetEnvVariable("CONTRIBUTOR_PRIVATE_ADDRESS"),
		client)

	manager.GetDataFromBlock(manager.GetHeader())

	PrintBalances(manager, provider, contributor)

	fmt.Println("deploying contract  ... ")
	campaign.Deploy(manager)
	PrintBalances(manager, provider, contributor)

	fmt.Println("contribute ...")
	campaign.Contribute(contributor)
	PrintBalances(manager, provider, contributor)

	campaign.IsPaid(contributor)

	campaign.GetBalance()

	request := entity.Request{
		ProviderAddress: provider.PublicAddress,
		Value:           big.NewInt(100),
		Description:     "water",
	}

	fmt.Println("create request ...")
	campaign.AddRequest(request, manager)
	PrintBalances(manager, provider, contributor)

	campaign.GetNumberOfRequests()

	campaign.ApproveRequest(contributor, big.NewInt(0))

	campaign.GetBalance()

	campaign.FinalizeRequest(manager, big.NewInt(0))

}

func PrintBalances(manager *account.Account, provider *account.Account, contributor *account.Account) {
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>Balance>>>>>>>>>>>>>>>>")
	fmt.Println("balance manager :", manager.GetBalance())
	fmt.Println("balance provider :", provider.GetBalance())
	fmt.Println("balance contributor :", contributor.GetBalance())
	fmt.Println("<<<<<<<<<<<<<<<<<<<<<<Balance<<<<<<<<<<<<<<<<<")

}
