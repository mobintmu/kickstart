package main

import (
	"kickstart/client"
	"kickstart/helper"

	"github.com/joho/godotenv"
)

func main() {

	// Load the .env file in the current directory
	godotenv.Load()

	client := client.Client{}

	address := helper.GetEnvVariable("ADDRESS")

	client.NewClient(address)

}
