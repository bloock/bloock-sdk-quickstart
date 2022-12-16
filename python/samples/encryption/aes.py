from bloock.client.entity.decrypter import AesDecrypter
from bloock.client.entity.encrypter import AesEncrypter
from utils.logger import Logger
from utils.sample import Sample
from utils.config import Config

from bloock.client.builder import RecordBuilder


def aes(_: Config):
    payload = "This will be encrypted"
    password = "a STRONG password"

    Logger.warn(f'The following payload will be encrypted: "{payload}"')

    # To encrypt we add an encrypter to the builder
    encrypted_record = (
        RecordBuilder.from_string(payload)
        .with_encrypter(AesEncrypter(password))
        .build()
    )

    Logger.success("Encryption successful")
    Logger.success(f"Encrypted payload: {encrypted_record.retrieve().decode()}")

    Logger.warn("Trying to decrypt with the valid password")

    # To decrypt we build a record from the encrypted record and add a decrypter
    decrypted_record = (
        RecordBuilder.from_record(encrypted_record)
        .with_decrypter(AesDecrypter(password))
        .build()
    )

    Logger.success(f"Decryption successful")

    hash = decrypted_record.get_hash()
    Logger.success(f"Hash: {hash}")
    Logger.success(f'Decrypted payload: "{decrypted_record.retrieve().decode()}"')

    Logger.warn("Trying to decrypt with invalid password")

    exception_was_thrown = False
    try:
        RecordBuilder.from_record(encrypted_record).with_decrypter(
            AesDecrypter("an invalid password")
        ).build()
    except:
        exception_was_thrown = True

    if not exception_was_thrown:
        raise Exception("The password was invalid so an error should've been returned!")

    Logger.success("The password was invalid so the record could not be decrypted")


Sample("AES encryption", aes)
