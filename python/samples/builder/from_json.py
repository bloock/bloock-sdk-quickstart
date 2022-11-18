from utils.logger import Logger
from utils.sample import Sample
from utils.config import Config

from bloock.client.builder import RecordBuilder

def from_json(_: Config):
    record = RecordBuilder.from_json('{"hello":"world"}').build()
    Logger.success('Record was created successfully')

    hash = record.get_hash()
    if hash != "586e9b1e1681ba3ebad5ff5e6f673d3e3aa129fcdb76f92083dbc386cdde4312":
        raise Exception("Unexpected hash received")

    Logger.success(f'Hash: {hash}')

Sample("builder_from_json", from_json)
