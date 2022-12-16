from bloock.client.client import Client
from bloock.client.entity.decrypter import EciesDecrypter
from bloock.client.entity.encrypter import EciesEncrypter
from utils.logger import Logger
from utils.sample import Sample
from utils.config import Config

from bloock.client.builder import RecordBuilder


def ecies(_: Config):
    sdk = Client()

    payload = "This will be encrypted"

    keypair = sdk.generate_ecies_keypair()

    Logger.warn(f'The following payload will be encrypted: "{payload}"')

    # To encrypt we add an encrypter to the builder
    encrypted_record = (
        RecordBuilder.from_string(payload)
        .with_encrypter(EciesEncrypter(keypair.public_key))
        .build()
    )

    Logger.success("Encryption successful")
    Logger.success(f"Encrypted payload: {encrypted_record.retrieve().decode()}")

    Logger.warn("Trying to decrypt with the public key")

    # To decrypt we build a record from the encrypted record and add a decrypter
    decrypted_record = (
        RecordBuilder.from_record(encrypted_record)
        .with_decrypter(EciesDecrypter(keypair.private_key))
        .build()
    )

    Logger.success(f"Decryption successful")

    hash = decrypted_record.get_hash()
    Logger.success(f"Hash: {hash}")
    Logger.success(f'Decrypted payload: "{decrypted_record.retrieve().decode()}"')

    Logger.warn("Trying to decrypt with invalid private key")

    exception_was_thrown = False
    try:
        RecordBuilder.from_record(encrypted_record).with_decrypter(
            EciesDecrypter("an invalid key")
        ).build()
    except:
        exception_was_thrown = True

    if not exception_was_thrown:
        raise Exception("The key was invalid so an error should've been returned!")

    Logger.success("The key was invalid so the record could not be decrypted")


Sample("ECIES encryption", ecies)
