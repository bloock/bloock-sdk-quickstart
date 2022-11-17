package main

import (
	"fmt"

	"github.com/bloock/bloock-sdk-go/v2"
	"github.com/bloock/bloock-sdk-go/v2/builder"
	"github.com/bloock/bloock-sdk-go/v2/client"
	"github.com/bloock/bloock-sdk-go/v2/client/entity"
	"github.com/bloock/bloock-sdk-quickstart/utils"
)

func main() {
	utils.Sample("Verification", func(c utils.Config) error {
		bloock.ApiKey = c.ApiKey
		sdk := client.NewClient()

		records := []string{}

		record, err := builder.NewRecordBuilderFromString("Hello world").Build()
		if err != nil {
			return err
		}

		hash, err := record.GetHash()
		if err != nil {
			return err
		}

		records = append(records, hash)
		receipt, err := sdk.SendRecords(records)
		if err != nil {
			return err
		}

		_, err = sdk.WaitAnchor(receipt[0].Anchor, entity.NewAnchorParams())
		if err != nil {
			return err
		}

		network := entity.NewNetworkParams()
		// we can specify the network we verify against or leave the default
		network.Network = entity.ListOfNetworks().BloockChain
		// we then verify the records and we will recive a timestamp
		// greater than 0 if the verification was successful
		timestamp, _ := sdk.VerifyRecords(records, network)
		fmt.Println(timestamp)
		return nil
	})

	utils.Sample("Verification Long", func(c utils.Config) error {
		bloock.ApiKey = c.ApiKey
		sdk := client.NewClient()

		records := []string{}

		record, err := builder.NewRecordBuilderFromString("Hello world").Build()
		if err != nil {
			return err
		}

		hash, err := record.GetHash()
		if err != nil {
			return err
		}

		records = append(records, hash)
		receipt, err := sdk.SendRecords(records)
		if err != nil {
			return err
		}

		_, err = sdk.WaitAnchor(receipt[0].Anchor, entity.NewAnchorParams())
		if err != nil {
			return err
		}

		network := entity.NewNetworkParams()
		// we can specify the network we verify against or leave the default
		network.Network = entity.ListOfNetworks().BloockChain

		proof, err := sdk.GetProof(records)
		if err != nil {
			return err
		}

		root, err := sdk.VerifyProof(proof)
		if err != nil {
			return err
		}

		// we then verify the root and we will recive a timestamp
		// greater than 0 if the verification was successful
		timestamp, err := sdk.ValidateRoot(root, network)
		if err != nil {
			return err
		}

		fmt.Println(timestamp)
		return nil
	})
}
