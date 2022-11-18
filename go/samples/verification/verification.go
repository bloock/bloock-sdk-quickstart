package main

import (
	"github.com/bloock/bloock-sdk-go/v2"
	"github.com/bloock/bloock-sdk-go/v2/builder"
	"github.com/bloock/bloock-sdk-go/v2/client"
	"github.com/bloock/bloock-sdk-go/v2/client/entity"
	"github.com/bloock/bloock-sdk-quickstart/utils"
	"github.com/fatih/color"
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

		params := entity.NewAnchorParams()
		// we can leave the params as default or we can specify a timeout
		params.Timeout = 120000 // default is 120000

		// Once we sent a record, we can wait for it's anochor
		color.Yellow("[+] Waiting for anchor...")
		anchor, err := sdk.WaitAnchor(receipt[0].Anchor, params)
		if err != nil {
			return err
		}
		color.Green("[✓] Anchor %+v done!", anchor)

		network := entity.NewNetworkParams()
		// we can specify the network we verify against or leave the default
		network.Network = entity.ListOfNetworks().BloockChain
		// we then verify the records and we will recive a timestamp
		// greater than 0 if the verification was successful
		timestamp, _ := sdk.VerifyRecords(records, network)
		color.Green("[✓] Timestamp: %d", timestamp)
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

		// Once we sent a record, we can wait for it's anochor
		color.Yellow("[+] Waiting for anchor...")
		anchor, err := sdk.WaitAnchor(receipt[0].Anchor, entity.NewAnchorParams())
		if err != nil {
			return err
		}
		color.Green("[✓] Anchor %d done!", anchor)

		// first we get the proof
		proof, err := sdk.GetProof(records)
		if err != nil {
			return err
		}

		// then verify it
		root, err := sdk.VerifyProof(proof)
		if err != nil {
			return err
		}

		network := entity.NewNetworkParams()
		// we can specify the network we verify against or leave the default (EthereumMainnet)
		network.Network = entity.ListOfNetworks().BloockChain
		// And finally validate the root
		timestamp, err := sdk.ValidateRoot(root, network)
		if err != nil {
			return err
		}

		// We will recive a timestamp greater than 0 if the validation was successful
		color.Green("[✓] Timestamp: %d", timestamp)
		return nil
	})
}
