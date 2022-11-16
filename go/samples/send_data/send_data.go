package main

import (
	"fmt"

	"github.com/bloock/bloock-sdk-go/v2"
	"github.com/bloock/bloock-sdk-go/v2/builder"
	"github.com/bloock/bloock-sdk-go/v2/client"
	"github.com/bloock/bloock-sdk-quickstart/utils"
)

func main() {
	utils.Sample("Send Data", func(c utils.Config) error {
		// we set the API key and create a client
		bloock.ApiKey = c.ApiKey
		sdk := client.NewClient()

		// we create an array of strings which will contain
		// the hashes of the records we want to send
		records := []string{}

		// first we create a record
		record, err := builder.NewRecordBuilderFromString("Hello world").Build()
		if err != nil {
			return err
		}

		// we then get its hash
		hash, err := record.GetHash()
		if err != nil {
			return err
		}

		// and append it to the array
		records = append(records, hash)

		// finally we can send the records
		receipt, err := sdk.SendRecords(records)
		if err != nil {
			return err
		}
		// we get a receipt with informationa about the transaction
		fmt.Printf("Record receipts: %+v", receipt)

		return nil
	})
}
