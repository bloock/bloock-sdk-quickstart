package main

import (
	"errors"
	"io/ioutil"

	"github.com/bloock/bloock-sdk-go/v2/builder"
	"github.com/bloock/bloock-sdk-quickstart/utils"
	"github.com/fatih/color"
)

func main() {
	utils.Sample("builder_from_file", func(c utils.Config) error {
		file, err := ioutil.ReadFile("resources/dummy.pdf")
		if err != nil {
			return err
		}

		record, err := builder.NewRecordBuilderFromFile(file).Build()
		if err != nil {
			return err
		}

		color.Green("[✓] Record was created successfully")

		hash, err := record.GetHash()
		if err != nil {
			return err
		}

		if hash != "43fa336d95e1634fee2d3031adc44ed9464472b94511af1facb87a1fee0b7e0e" {
			return errors.New("Unexpected hash received")
		}

		color.Green("[✓] %s", hash)

		// we can get the file back if needed
		err = ioutil.WriteFile("resources/output.pdf", record.Retrieve(), 0644)
		if err != nil {
			return err
		}

		return nil
	})
}
