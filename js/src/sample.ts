import { workerData } from 'node:worker_threads'
import { Config } from './config'
import * as colors from 'colors'

export class Sample {
    public static run(name: string, fn: (config: Config) => Promise<any>) {
        let config = workerData as Config;

        console.log(colors.yellow(`${name}: Started`))
        fn(config).then(() => {
            console.log(colors.green(`${name}: Successful`))
        }).catch(err => {
            console.log(colors.red(`${name}: Failure`))
            console.log(colors.red(`${name}: ${err}`))
        })
    }
}