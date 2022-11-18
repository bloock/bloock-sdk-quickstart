import bloock
from bloock.client.client import Client, Network
from bloock.client.builder import RecordBuilder

from utils.logger import Logger
from utils.sample import Sample
from utils.config import Config

def verification(c: Config):
    bloock.api_key = c.api_key
    client = Client()

    record = RecordBuilder.from_string("Hello world").build()
    hash = record.get_hash()
    records = [hash]
    
    receipts = client.send_records(records)
    
    Logger.info("Waiting for anchor...")
    # we can optionally specify a timeout (if not set, default is 120000) 
    anchor = client.wait_anchor(receipts[0].anchor, timeout=120000)
    Logger.success(f'Anchor {anchor.__dict__} done!')

    # we can optionally specify a network (if not set, default is Ethereum Mainnet)
    timestamp = client.verify_records(records, Network.ETHEREUM_MAINNET)
    Logger.success(f'Timestamp: {timestamp}')

def verification_long(c: Config):
    bloock.api_key = c.api_key
    client = Client()

    record = RecordBuilder.from_string("Hello world").build()
    hash = record.get_hash()
    records = [hash]
    
    receipts = client.send_records(records)
    
    Logger.info("Waiting for anchor...")
    # we can optionally specify a timeout (if not set, default is 120000) 
    anchor = client.wait_anchor(receipts[0].anchor, timeout=120000)
    Logger.success(f'Anchor {anchor.__dict__} done!')

    # first we get the proof
    proof = client.get_proof(records);

    # then verify it
    root = client.verify_proof(proof);

    # And finally validate the root. We can optionally specify a network 
    # (if not set, default is Ethereum Mainnet)
    timestamp = client.validate_root(root, Network.ETHEREUM_MAINNET)

    # We will recive a timestamp greater than 0 if the validation was successful
    Logger.success(f'Timestamp: {timestamp}');

Sample("verification", verification)
Sample("verification_long", verification_long)
