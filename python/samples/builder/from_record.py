from utils.sample import Sample
from utils.config import Config
from colorama import Fore, Style

from bloock.client.builder import RecordBuilder

def from_record(_: Config):
    record = RecordBuilder.from_string("Hello world").build()

    record = RecordBuilder.from_record(record).build()

    print(Fore.GREEN + f'[✓] Record was created successfully' + Style.RESET_ALL)

    hash = record.get_hash()
    if hash != "ed6c11b0b5b808960df26f5bfc471d04c1995b0ffd2055925ad1be28d6baadfd":
        raise Exception("Unexpected hash received")

    print(Fore.GREEN + f'[✓] Hash: {hash}' + Style.RESET_ALL)

Sample("builder_from_record", from_record)
