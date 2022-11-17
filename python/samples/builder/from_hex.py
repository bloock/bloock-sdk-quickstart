from utils.sample import Sample
from utils.config import Config
from colorama import Fore, Style

from bloock.client.builder import RecordBuilder

def from_hex(_: Config):
    record = RecordBuilder.from_hex("1234567890abcdef").build()
    print(Fore.GREEN + f'[✓] Record was created successfully' + Style.RESET_ALL)

    hash = record.get_hash()
    if hash != "ed8ab4fde4c4e2749641d9d89de3d920f9845e086abd71e6921319f41f0e784f":
        raise Exception("Unexpected hash received")

    print(Fore.GREEN + f'[✓] Hash: {hash}' + Style.RESET_ALL)

Sample("builder_from_hex", from_hex)
