package main

import (
	"errors"

	"github.com/bloock/bloock-sdk-go/v2/builder"
	"github.com/bloock/bloock-sdk-quickstart/utils"
	"github.com/bloock/bloock-sdk-quickstart/utils/logger"
)

func main() {
	utils.Sample("builder_from_hex", func(c utils.Config) error {
		record, err := builder.NewRecordBuilderFromHex("1234567890abcdef").Build()
		if err != nil {
			return err
		}

		logger.Success("Record was created successfully")

		hash, err := record.GetHash()
		if err != nil {
			return err
		}

		if hash != "ed8ab4fde4c4e2749641d9d89de3d920f9845e086abd71e6921319f41f0e784f" {
			return errors.New("Unexpected hash received")
		}

		logger.Success("Hash: " + hash)

		return nil
	})
}
