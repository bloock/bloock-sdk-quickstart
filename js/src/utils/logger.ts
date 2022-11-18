import * as colors from 'colors'

export class Logger {
    public static success(message: string) {
        console.log(colors.green(`[✓] ${message}`));
    }

    public static warn(message: string) {
        console.log(colors.yellow(`[!] ${message}`));
    }

    public static info(message: string) {
        console.log(colors.yellow(`[+] ${message}`));
    }

    public static error(message: string) {
        console.log(colors.red(`[x] ${message}`));
    }
}
