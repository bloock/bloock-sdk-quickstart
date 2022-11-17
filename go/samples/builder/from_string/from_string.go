package main

import (
	"errors"

	"github.com/bloock/bloock-sdk-go/v2/builder"
	"github.com/bloock/bloock-sdk-quickstart/utils"
	"github.com/fatih/color"
)

func main() {
	utils.Sample("builder_from_string", func(c utils.Config) error {
		record, err := builder.NewRecordBuilderFromString("Hello world").Build()
		if err != nil {
			return err
		}

		color.Green("[✓]  Record was created successfully")

		hash, err := record.GetHash()
		if err != nil {
			return err
		}

		if hash != "ed6c11b0b5b808960df26f5bfc471d04c1995b0ffd2055925ad1be28d6baadfd" {
			return errors.New("Unexpected hash received")
		}

		color.Green("[✓]  Hash: %s", hash)

		return nil
	})
}
