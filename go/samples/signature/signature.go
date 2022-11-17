package main

import (
	"github.com/bloock/bloock-sdk-go/v2"
	"github.com/bloock/bloock-sdk-go/v2/builder"
	"github.com/bloock/bloock-sdk-go/v2/client"
	"github.com/bloock/bloock-sdk-go/v2/client/entity"
	"github.com/bloock/bloock-sdk-quickstart/utils"
	"github.com/fatih/color"
)

func main() {
	utils.Sample("signature", func(c utils.Config) error {
        bloock.ApiKey = c.ApiKey
        sdk := client.NewClient()

        payload := "This will be encrypted"

        color.Yellow("[+] The following payload will be encrypted: '%s'\n", payload)

        keys, err := sdk.GenerateKeys()
        if err != nil {
            return err
        }

		signedRecord, err := builder.NewRecordBuilderFromString(payload).
			WithSigner(entity.NewEcsdaSigner(keys.PrivateKey)).
			Build()

        if err != nil {
            return err
        }

        color.Green("[✓] Record was signed succesfully")

        // we can add another signature with a different key 
        keys, err = sdk.GenerateKeys()
        if err != nil {
            return err
        }

		signedRecord, err = builder.NewRecordBuilderFromRecord(signedRecord).
			WithSigner(entity.NewEcsdaSigner(keys.PrivateKey)).
			Build()

        if err != nil {
            return err
        }

        color.Green("[✓] Record was signed succesfully")

        hash, err := signedRecord.GetHash()
        if err != nil {
            return err
        }
        color.Green("[✓] Hash: %s", hash)

        signatures, err := signedRecord.GetSignatures()
        if err != nil {
            return err
        }

        for i, signature := range signatures {
            color.Green("[✓] Signature %d: %+v", i+1, signature.Signature)
        }

        return nil
    })
}
