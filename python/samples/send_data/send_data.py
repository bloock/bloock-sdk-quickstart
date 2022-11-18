import bloock
from bloock.client.client import Client
from bloock.client.builder import RecordBuilder

from utils.logger import Logger
from utils.sample import Sample
from utils.config import Config

def send_data(c: Config):
    bloock.api_key = c.api_key
    client = Client()

    # we create an array of strings which will contain
    # the hashes of the records we want to send
    records = []
    record = RecordBuilder.from_string("Hello world").build()
    # we can get the hash of the record
    hash = record.get_hash()
    
    # and append it to the array
    records.append(hash)
    
    # finally we can send the records
    send_receipts = client.send_records(records)
    
    # we get a receipt with informationa about the transaction
    Logger.success(f'Record receipts: {list(map(lambda x: x.__dict__, send_receipts))}')

Sample("send_data", send_data)
