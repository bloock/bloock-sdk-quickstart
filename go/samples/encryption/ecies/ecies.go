package main

import (
	"errors"

	"github.com/bloock/bloock-sdk-go/v2/builder"
	"github.com/bloock/bloock-sdk-go/v2/client"
	"github.com/bloock/bloock-sdk-go/v2/client/entity"
	"github.com/bloock/bloock-sdk-quickstart/utils"
	"github.com/bloock/bloock-sdk-quickstart/utils/logger"
)

func main() {
	utils.Sample("ECIES encryption/decryption", func(c utils.Config) error {
		payload := "This will be encrypted"
		sdk := client.NewClient()

        keypair, err := sdk.GenerateEciesKeyPair()

		if err != nil {
			return err
		}

		logger.Info("The following payload will be encrypted: " + payload)

		// To encrypt we add an encrypter to the builder
		encryptedRecord, err := builder.NewRecordBuilderFromString(payload).
			WithEncrypter(entity.NewEciesEncrypter(keypair.PublicKey)).
			Build()

		if err != nil {
			return err
		}

		logger.Success("Encryption successful")

		logger.Success("Encrypted payload: " + string(encryptedRecord.Retrieve()))

		logger.Info("Trying to decrypt with the private key")

		// To decrypt we build a record from the encrypted record and add a decrypter
		decryptedRecord, err := builder.NewRecordBuilderFromRecord(encryptedRecord).
			WithDecrypter(entity.NewEciesDecrypter(keypair.PrivateKey)).
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

		logger.Info("Trying to decrypt with invalid private key")

		_, err = builder.NewRecordBuilderFromRecord(encryptedRecord).
			WithDecrypter(entity.NewEciesDecrypter("an invalid key")).
			Build()

		if err == nil {
			return errors.New("The key was invalid so an error should've been returned!")
		}

		logger.Success("The key was invalid so the record could not be decrypted")

		return nil
	})
}
