from bloock.client.builder import RecordBuilder
from bloock.client.client import Client

from bloock.client.entity.proof import AnchorNetwork, Proof, ProofAnchor
from utils.logger import Logger

from utils.sample import Sample
from utils.config import Config


def set_proof(_: Config):
    record = RecordBuilder.from_string("Hello world").build()
    # we can set the proof of a record
    record.set_proof(Proof(
            leaves=["b20422fcfbe5b818ee305b44d5aaf427d103f8c5791b79c772ce82c747b2cd0d"],
            nodes=["b20422fcfbe5b818ee305b44d5aaf427d103f8c5791b79c772ce82c747b2cd0d"],
            depth="10101010",
            bitmap="01010101",
            anchor=ProofAnchor(
                anchor_id=10,
                networks=[
                    AnchorNetwork(
                        name="Ethereum",
                        state="state",
                        tx_hash="b20422fcfbe5b818ee305b44d5aaf427d103f8c5791b79c772ce82c747b2cd0d",
                    )
                ],
                root="b20422fcfbe5b818ee305b44d5aaf427d103f8c5791b79c772ce82c747b2cd0d",
                status="success",
            ),
        ))

    proof = Client().get_proof([record])
    Logger.success(f"The following proof was set: {proof.__dict__}")


Sample("set_proof", set_proof)
