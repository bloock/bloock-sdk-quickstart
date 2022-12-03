import os
import sys
from utils.config import Config

from utils.logger import Logger


def Sample(name, fn):
    config = Config(api_key=os.getenv("API_KEY"))

    Logger.info(f"{name}: Started")
    try:
        fn(config)
    except Exception as e:
        Logger.err("{name}: Failure")
        Logger.err(f"{name}: {e}")
        sys.exit(1)
    else:
        Logger.success(f"{name}: Successful")
