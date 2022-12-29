package main

import (
	"errors"

	"github.com/bloock/bloock-sdk-go/v2"
	"github.com/bloock/bloock-sdk-go/v2/builder"
	"github.com/bloock/bloock-sdk-go/v2/client/entity"
	"github.com/bloock/bloock-sdk-quickstart/utils"
	"github.com/bloock/bloock-sdk-quickstart/utils/logger"
)

func main() {
	utils.Sample("builder_from_hosted_loader", func(c utils.Config) error {
		record, err := builder.NewRecordBuilderFromString("Hello world").Build()
		if err != nil {
			return err
		}

		hash, err := record.GetHash()
		if err != nil {
			return err
		}

		if hash != "ed6c11b0b5b808960df26f5bfc471d04c1995b0ffd2055925ad1be28d6baadfd" {
			return errors.New("Unexpected hash received")
		}

		bloock.ApiKey = c.ApiKey

		result, err := record.Publish(entity.NewHostedPublisher())
		if err != nil {
			return err
		}

		if result != hash {
			return errors.New("Publish's result should be equal to the record's hash")
		}

		record, err = builder.NewRecordBuilderFromLoader(entity.NewHostedLoader(result)).Build()
		if err != nil {
			return err
		}

		logger.Success("Record was created successfully")

		hash, err = record.GetHash()
		if err != nil {
			return err
		}

		if hash != result {
			return errors.New("Unexpected hash received")
		}

		logger.Success("Hash: " + hash)

		return nil
	})

	utils.Sample("builder_from_ipfs_loader", func(c utils.Config) error {
		record, err := builder.NewRecordBuilderFromString("Hello world").Build()
		if err != nil {
			return err
		}

		hash, err := record.GetHash()
		if err != nil {
			return err
		}

		if hash != "ed6c11b0b5b808960df26f5bfc471d04c1995b0ffd2055925ad1be28d6baadfd" {
			return errors.New("Unexpected hash received")
		}

		bloock.ApiKey = c.ApiKey

		result, err := record.Publish(entity.NewIpfsPublisher())
		if err != nil {
			return err
		}

		if result != hash {
			return errors.New("Publish's result should be equal to the record's hash")
		}

		record, err = builder.NewRecordBuilderFromLoader(entity.NewIpfsLoader(result)).Build()
		if err != nil {
			return err
		}

		logger.Success("Record was created successfully")

		hash, err = record.GetHash()
		if err != nil {
			return err
		}

		if hash != result {
			return errors.New("Unexpected hash received")
		}

		logger.Success("Hash: " + hash)

		return nil
	})
}
