from utils.logger import Logger
from utils.sample import Sample
from utils.config import Config

from bloock.client.builder import RecordBuilder


def from_bytes(_: Config):
    record = RecordBuilder.from_bytes(bytes([1, 2, 3, 4, 5])).build()
    Logger.success("Record was created successfully")

    hash = record.get_hash()
    if hash != "7d87c5ea75f7378bb701e404c50639161af3eff66293e9f375b5f17eb50476f4":
        raise Exception("Unexpected hash received")

    Logger.success(f"Hash: {hash}")


Sample("builder_from_bytes", from_bytes)
