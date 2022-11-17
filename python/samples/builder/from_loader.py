from bloock.client.entity.loader import HostedLoader
from bloock.client.entity.publisher import HostedPublisher
from utils.sample import Sample
from utils.config import Config
from colorama import Fore, Style

from bloock.client.builder import RecordBuilder

def from_loader(_: Config):
    record = RecordBuilder.from_string("Hello world").build()

    hash = record.get_hash()
    if hash != "ed6c11b0b5b808960df26f5bfc471d04c1995b0ffd2055925ad1be28d6baadfd":
        raise Exception("Unexpected hash received")

    result = record.publish(HostedPublisher())
    if result != hash:
        raise Exception("Publish's result should be equal to the record's hash")

    record = RecordBuilder.from_loader(HostedLoader(hash=result)).build()
    print(Fore.GREEN + f'[✓] Record was created successfully' + Style.RESET_ALL)


    print(Fore.GREEN + f'[✓] Hash: {hash}' + Style.RESET_ALL)

Sample("builder_from_loader", from_loader)
