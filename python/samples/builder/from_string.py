from utils.sample import Sample
from utils.config import Config

def from_string(config: Config):
    print("from_string")
    raise Exception("an error")

Sample("builder_from_string", from_string)
