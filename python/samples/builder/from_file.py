from utils.logger import Logger
from utils.sample import Sample
from utils.config import Config

from bloock.client.builder import RecordBuilder

def from_file(_: Config):
    data = bytes()
    # we can read a file as an array of bytes
    with open("sample.pdf", "rb") as file:
        data = bytes(file.read())

    # and build a record from it
    record = RecordBuilder.from_file(data).build()
    
    Logger.success('Record was created successfully')

    hash = record.get_hash()
    if hash != "43fa336d95e1634fee2d3031adc44ed9464472b94511af1facb87a1fee0b7e0e":
        raise Exception("Unexpected hash received")

    Logger.success(f'Hash: {hash}')

    # we can get the file back if needed
    with open("output.pdf", "wb") as file:
        file.write(record.retrieve())

Sample("builder_from_file", from_file)
