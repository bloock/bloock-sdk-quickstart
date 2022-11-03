import * as glob from 'glob';
import * as swc from '@swc/core';
import { isMainThread, Worker, WorkerOptions } from 'node:worker_threads';
import { Config } from './config';

export class Runner {
    public static async run(config: Config) {
        if (isMainThread) {
            const samples = glob.sync('src/samples/**/*.ts', { absolute: true });
            for (const s of samples) {
                runWorker(s, config)
            }
        }
    }
}

const runWorker = (file: string, config: Config) => {
    let wkOpts: WorkerOptions = {
        eval: true,
        workerData: {
            __filename: file,
            ...config
        }
    }
    return new Worker(`
            const wk = require('worker_threads');
            require('ts-node').register();
            let file = wk.workerData.__filename;
            delete wk.workerData.__filename;
            require(file);
        `,
        wkOpts
    );
}
