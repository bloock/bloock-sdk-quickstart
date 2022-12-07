package main

import (
	"github.com/bloock/bloock-sdk-go/v2/client"
	"github.com/bloock/bloock-sdk-go/v2/client/entity"
	"github.com/bloock/bloock-sdk-quickstart/utils"
)

func main() {
	utils.Sample("set provider", func(c utils.Config) error {
		// we can change the provider by specifing the network and the new url of the provider
		client.SetProvider(entity.ListOfNetworks().EthereumMainnet, "https://eth.someprovider.com")
		return nil
	})

	utils.Sample("set contract address", func(c utils.Config) error {
		// we can change the contract address by specifing the network and the new contract address
		client.SetContractAddess(entity.ListOfNetworks().EthereumMainnet, "some contract address")
		return nil
	})
}
