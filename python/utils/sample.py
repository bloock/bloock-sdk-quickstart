import os
from utils.config import Config
from colorama import Fore, Style

def Sample(name, fn):
    config = Config(api_key=os.getenv("API_KEY"))

    print(Fore.YELLOW + f'[+] {name}: Started' + Style.RESET_ALL)
    try:
        fn(config)
    except Exception as e:
        print(Fore.RED + f'[x] {name}: Failure' + Style.RESET_ALL)
        print(Fore.RED + f'[x] {name}: {e}' + Style.RESET_ALL)
    else:
        print(Fore.GREEN + f'[✓] {name}: Successful' + Style.RESET_ALL)

