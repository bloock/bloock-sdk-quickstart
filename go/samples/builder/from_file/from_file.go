package main

import (
	"errors"
	"io/ioutil"

	"github.com/bloock/bloock-sdk-go/v2/builder"
	"github.com/bloock/bloock-sdk-quickstart/utils"
	"github.com/fatih/color"
)

func main() {
	utils.Sample("builder_from_bytes", func(c utils.Config) error {
        file, err := ioutil.ReadFile("dummy.pdf")
        if err != nil {
            return err
        }

        record, err := builder.NewRecordBuilderFromFile(file).Build()
        if err != nil {
            return err
        }

        color.Green("=> Record was created successfully")

        hash, err := record.GetHash()
        if err != nil {
            return err
        }
        
        if hash != "TODO" {
            return errors.New("Unexpected hash received")
        }

        color.Green("[✓] %s", hash)

		return nil
	})
}
