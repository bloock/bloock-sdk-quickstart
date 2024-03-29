import bloock
from bloock.client.client import Client
from bloock.client.builder import RecordBuilder
from bloock.client.entity.signer import EcsdaSigner

from utils.logger import Logger
from utils.sample import Sample
from utils.config import Config


def signature(c: Config):
    bloock.api_key = c.api_key
    client = Client()

    keys = client.generate_keys()

    signed_record = (
        RecordBuilder.from_string("Hello world")
        .with_signer(EcsdaSigner(keys.private_key))
        .build()
    )

    Logger.success("Record was signed sucessfully")

    # we can add another signature with a different key
    keys = client.generate_keys()

    Logger.info("Adding another signature with a common name")

    name = "Some name"

    signed_record = (
        RecordBuilder.from_record(signed_record)
        .with_signer(EcsdaSigner(keys.private_key, common_name=name))
        .build()
    )

    Logger.success("Record was signed sucessfully")

    hash = signed_record.get_hash()

    Logger.success(f"Hash: {hash}")
    if hash != "ed6c11b0b5b808960df26f5bfc471d04c1995b0ffd2055925ad1be28d6baadfd":
        raise Exception("Unexpected hash received")

    signatures = signed_record.get_signatures()

    for i, signature in enumerate(signatures):
        Logger.success(f"Signature {i + 1}: {signature.signature}")

    retrieved_name = signatures[1].get_common_name()

    if retrieved_name != name:
        raise Exception("The retrieved name is not the expected name")

    Logger.success(f"Common name for signature is {retrieved_name}")


Sample("signature", signature)
