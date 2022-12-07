import bloock

from bloock.client.entity.network import Network

from utils.sample import Sample
from utils.config import Config


def set_provider(_: Config):
    # we can change the provider by specifing the network and the new url of the provider
    bloock.set_provider(Network.ETHEREUM_MAINNET, "https://eth.someprovider.com")


def set_contract_address(_: Config):
    # we can change the contract address by specifing the network and the new contract address
    bloock.set_contract_address(Network.ETHEREUM_MAINNET, "some contract address")


Sample("set_provider", set_provider)
Sample("set_contract_address", set_contract_address)
