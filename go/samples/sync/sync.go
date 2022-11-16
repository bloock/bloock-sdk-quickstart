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
	utils.Sample("Sync", func(c utils.Config) error {
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
		// we can leve the params as default or we can specify a timeout
		params.Timeout = 60000 // default is 120000

		// Once we sent a record, we can wait for it's anochor
		fmt.Println("Waiting for anchor...")
		anchor, err := sdk.WaitAnchor(receipt[0].Anchor, params)
		if err != nil { 
            return err
        }
		fmt.Println("Done!")

        fmt.Println(anchor)

        return nil
	})
}
