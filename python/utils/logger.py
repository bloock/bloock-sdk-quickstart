from colorama import Fore, Style


class Logger:
    @staticmethod
    def success(message: str):
        print(Fore.GREEN + "[✓] " + message + Style.RESET_ALL)

    @staticmethod
    def warn(message: str):
        print(Fore.YELLOW + "[!] " + message + Style.RESET_ALL)

    @staticmethod
    def info(message: str):
        print(Fore.CYAN + "[i] " + message + Style.RESET_ALL)

    @staticmethod
    def err(message: str):
        print(Fore.RED + "[x] " + message + Style.RESET_ALL)
