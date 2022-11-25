package main

import (
	"errors"

	"github.com/bloock/bloock-sdk-go/v2/builder"
	"github.com/bloock/bloock-sdk-quickstart/utils"
	"github.com/bloock/bloock-sdk-quickstart/utils/logger"
)

func main() {
	utils.Sample("builder_from_bytes", func(c utils.Config) error {
		record, err := builder.NewRecordBuilderFromJSON("{\"hello\":\"world\"}").Build()
		if err != nil {
			return err
		}

		logger.Success("Record was created successfully")

		hash, err := record.GetHash()
		if err != nil {
			return err
		}

		if hash != "586e9b1e1681ba3ebad5ff5e6f673d3e3aa129fcdb76f92083dbc386cdde4312" {
			return errors.New("Unexpected hash received")
		}

		logger.Success("Hash: " + hash)

		return nil
	})
}
