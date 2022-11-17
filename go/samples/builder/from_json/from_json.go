package main

import (
	"errors"

	"github.com/bloock/bloock-sdk-go/v2/builder"
	"github.com/bloock/bloock-sdk-quickstart/utils"
	"github.com/fatih/color"
)

func main() {
	utils.Sample("builder_from_bytes", func(c utils.Config) error {
		record, err := builder.NewRecordBuilderFromJSON("{\"hello\":\"world\"}").Build()
		if err != nil {
			return err
		}

		color.Green("[✓]  Record was created successfully")

		hash, err := record.GetHash()
		if err != nil {
			return err
		}

		if hash != "586e9b1e1681ba3ebad5ff5e6f673d3e3aa129fcdb76f92083dbc386cdde4312" {
			return errors.New("Unexpected hash received")
		}

		color.Green("[✓]  Hash: %s", hash)

		return nil
	})
}
