from bloock import bloock
from bloock.client.entity.loader import HostedLoader
from bloock.client.entity.publisher import HostedPublisher
from utils.logger import Logger
from utils.sample import Sample
from utils.config import Config

from bloock.client.builder import RecordBuilder


def from_loader(config: Config):
    record = RecordBuilder.from_string("Hello world").build()

    hash = record.get_hash()
    if hash != "ed6c11b0b5b808960df26f5bfc471d04c1995b0ffd2055925ad1be28d6baadfd":
        raise Exception("Unexpected hash received")

    bloock.api_key = config.api_key

    result = record.publish(HostedPublisher())
    if result != hash:
        raise Exception("Publish's result should be equal to the record's hash")

    record = RecordBuilder.from_loader(HostedLoader(hash=result)).build()
    Logger.success("Record was created successfully")

    Logger.success(f"Hash: {hash}")


Sample("builder_from_loader", from_loader)
