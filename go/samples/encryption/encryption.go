package main

import (
	"errors"

	"github.com/bloock/bloock-sdk-go/v2/builder"
	"github.com/bloock/bloock-sdk-go/v2/client/entity"
	"github.com/bloock/bloock-sdk-quickstart/utils"
	"github.com/bloock/bloock-sdk-quickstart/utils/logger"
)

func main() {
	utils.Sample("simple encryption/decryption", func(c utils.Config) error {
		payload := "This will be encrypted"
		password := "a STRONG password"

		logger.Info("The following payload will be encrypted: " + payload)

		// To encrypt we add an encrypter to the builder
		encryptedRecord, err := builder.NewRecordBuilderFromString(payload).
			WithEncrypter(entity.NewAesEncrypter(password)).
			Build()

		if err != nil {
			return err
		}

		logger.Success("Encryption successful")

		logger.Success("Encrypted payload: " + string(encryptedRecord.Retrieve()))

		logger.Info("Trying to decrypt with the valid password")

		// To decrypt we build a record from the encrypted record and add a decrypter
		decryptedRecord, err := builder.NewRecordBuilderFromRecord(encryptedRecord).
			WithDecrypter(entity.NewAesDecrypter(password)).
			Build()

		if err != nil {
			return err
		}

		logger.Success("Decryption successful")

		hash, err := decryptedRecord.GetHash()
		if err != nil {
			return err
		}
		logger.Success("Hash: " + hash)

		logger.Success("Decrypted payload: " + string(decryptedRecord.Retrieve()))

		logger.Info("Trying to decrypt with invalid password")

		_, err = builder.NewRecordBuilderFromRecord(encryptedRecord).
			WithDecrypter(entity.NewAesDecrypter("an invalid password")).
			Build()

		if err == nil {
			return errors.New("The password was invalid so an error should've been returned!")
		}

		logger.Success("The password was invalid so the record could not be decrypted")

		return nil
	})
}
