import { Config } from "./config";
import { Logger } from "./logger";

export class Sample {
  public static run(name: string, fn: (config: Config) => Promise<any>) {
    let config: Config = {
      apiKey: process.env["API_KEY"],
      apiHost: process.env["API_HOST"]
    };

    Logger.info(`${name}: Started`);
    fn(config)
      .then(() => {
        Logger.success(`${name}: Successful`);
      })
      .catch(err => {
        Logger.error(`${name}: Failure`);
        Logger.error(`${name}: ${err}`);
        process.exit(1);
      });
  }
}
