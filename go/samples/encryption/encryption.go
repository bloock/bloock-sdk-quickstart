package main

import (
	"errors"

	"github.com/bloock/bloock-sdk-go/v2/builder"
	"github.com/bloock/bloock-sdk-go/v2/client/entity"
	"github.com/bloock/bloock-sdk-quickstart/utils"
	"github.com/fatih/color"
)

func main() {
	utils.Sample("simple encryption/decryption", func(c utils.Config) error {
        payload := "This will be encrypted"
        password := "a STRONG password here"

        color.Yellow("[+] The following payload will be encrypted: '%s'\n", payload)

		encryptedRecord, err := builder.NewRecordBuilderFromString(payload).
			WithEncrypter(entity.NewAesEncrypter(password)).
			Build()

        if err != nil {
            return err
        }

        color.Green("[✓] Encryption succesful")

        color.Yellow("[+] Trying to decrypt with the same password")

        decryptedRecord, err := builder.NewRecordBuilderFromRecord(encryptedRecord).
            WithDecrypter(entity.NewAesDecrypter(password)).
            Build()

        if err != nil {
            return err
        }

        color.Green("[✓] Decryption succesful")

        hash, err := decryptedRecord.GetHash()
        if err != nil {
            return err
        }
        color.Green("[✓] Hash: %s", hash)

        color.Yellow("[+] Trying to decrypt with invalid password")

        _, err = builder.NewRecordBuilderFromRecord(encryptedRecord).
            WithDecrypter(entity.NewAesDecrypter("an invalid password here")).
            Build()

        if err == nil {
            return errors.New("The password was invalid so an error should've been returned!")
        }

        color.Green("[✓] The password was invalid so the record could not be decrypted")

        return nil
    })
}
