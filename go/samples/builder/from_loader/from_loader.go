package main

import (
	"errors"

	"github.com/bloock/bloock-sdk-go/v2/builder"
	"github.com/bloock/bloock-sdk-go/v2/client/entity"
	"github.com/bloock/bloock-sdk-quickstart/utils"
	"github.com/fatih/color"
)

func main() {
	utils.Sample("builder_from_loader", func(c utils.Config) error {
		record, err := builder.NewRecordBuilderFromString("Hello world").Build()
        if err != nil {
            return err
        }

		hash, err := record.GetHash()
        if err != nil {
            return err
        }

		result, err := record.Publish(entity.NewHostedPublisher())
        if result != hash {
            return errors.New("Publish's result should be equal to the record's hash")
        }

		record, err = builder.NewRecordBuilderFromLoader(entity.NewHostedLoader(result)).Build()
        if err != nil {
            return err
        }

        color.Green("=> Record was created successfully")

        hash, err = record.GetHash()
        if err != nil {
            return err
        }
        
        if hash != result {
            return errors.New("Unexpected hash received")
        }

        color.Green("=> Hash: %s", hash)

		return nil
	})
}
