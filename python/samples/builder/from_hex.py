from utils.logger import Logger
from utils.sample import Sample
from utils.config import Config

from bloock.client.builder import RecordBuilder


def from_hex(_: Config):
    record = RecordBuilder.from_hex("1234567890abcdef").build()
    Logger.success("Record was created successfully")

    hash = record.get_hash()
    if hash != "ed8ab4fde4c4e2749641d9d89de3d920f9845e086abd71e6921319f41f0e784f":
        raise Exception("Unexpected hash received")

    Logger.success(f"Hash: {hash}")


Sample("builder_from_hex", from_hex)
