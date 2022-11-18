package main

import (
	"errors"

	"github.com/bloock/bloock-sdk-go/v2/builder"
	"github.com/bloock/bloock-sdk-quickstart/utils"
	"github.com/bloock/bloock-sdk-quickstart/utils/logger"
)

func main() {
	utils.Sample("builder_from_string", func(c utils.Config) error {
		record, err := builder.NewRecordBuilderFromString("Hello world").Build()
		if err != nil {
			return err
		}

        logger.Success("Record was created successfully")

		hash, err := record.GetHash()
		if err != nil {
			return err
		}

		if hash != "ed6c11b0b5b808960df26f5bfc471d04c1995b0ffd2055925ad1be28d6baadfd" {
			return errors.New("Unexpected hash received")
		}

        logger.Success("Hash: " + hash)

		return nil
	})
}
