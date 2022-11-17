from utils.sample import Sample
from utils.config import Config
from colorama import Fore, Style

from bloock.client.builder import RecordBuilder

def from_json(_: Config):
    record = RecordBuilder.from_json('{"hello":"world"}').build()
    print(Fore.GREEN + f'[✓] Record was created successfully' + Style.RESET_ALL)

    hash = record.get_hash()
    if hash != "586e9b1e1681ba3ebad5ff5e6f673d3e3aa129fcdb76f92083dbc386cdde4312":
        raise Exception("Unexpected hash received")

    print(Fore.GREEN + f'[✓] Hash: {hash}' + Style.RESET_ALL)

Sample("builder_from_json", from_json)
