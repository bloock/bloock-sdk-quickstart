package main

import (
	"errors"

	"github.com/bloock/bloock-sdk-go/v2/builder"
	"github.com/bloock/bloock-sdk-quickstart/utils"
	"github.com/fatih/color"
)

func main() {
	utils.Sample("builder_from_bytes", func(c utils.Config) error {
        record, err := builder.NewRecordBuilderFromBytes([]byte{1, 2, 3, 4, 5}).Build()
        if err != nil {
            return err
        }

        color.Green("=> Record was created successfully")

        hash, err := record.GetHash()
        if err != nil {
            return err
        }
        
        if hash != "7d87c5ea75f7378bb701e404c50639161af3eff66293e9f375b5f17eb50476f4" {
            return errors.New("Unexpected hash received")
        }

        color.Green("=> Hash: %s", hash)

		return nil
	})
}
