package main

import (
	"errors"
	"fmt"

	"github.com/bloock/bloock-sdk-go/v2"
	"github.com/bloock/bloock-sdk-go/v2/builder"
	"github.com/bloock/bloock-sdk-go/v2/client"
	"github.com/bloock/bloock-sdk-go/v2/client/entity"
	"github.com/bloock/bloock-sdk-quickstart/utils"
	"github.com/bloock/bloock-sdk-quickstart/utils/logger"
)

func main() {
	utils.Sample("signature", func(c utils.Config) error {
		bloock.ApiKey = c.ApiKey
		sdk := client.NewClient()

		keys, err := sdk.GenerateKeys()
		if err != nil {
			return err
		}

		signedRecord, err := builder.NewRecordBuilderFromString("Hello world").
			WithSigner(entity.NewEcsdaSigner(keys.PrivateKey)).
			Build()

		if err != nil {
			return err
		}

		logger.Success("Record was signed successfully")

		// we can add another signature with a different key
		keys, err = sdk.GenerateKeys()
		if err != nil {
			return err
		}

		logger.Info("Adding another signature")
		signedRecord, err = builder.NewRecordBuilderFromRecord(signedRecord).
			WithSigner(entity.NewEcsdaSigner(keys.PrivateKey)).
			Build()

		if err != nil {
			return err
		}

		logger.Success("Record was signed successfully")

		hash, err := signedRecord.GetHash()
		if err != nil {
			return err
		}

        if hash != "b6e6816e3c6180fcbda27048f033cf2b6f2a627864240c4c85558bcbece2a2e4" {
			return errors.New("Unexpected hash received")
        }

		logger.Success("Hash: " + hash)

		signatures, err := signedRecord.GetSignatures()
		if err != nil {
			return err
		}

		for i, signature := range signatures {
			logger.Success(fmt.Sprintf("Signature %d: %+v", i+1, signature.Signature))
		}

		return nil
	})
}
