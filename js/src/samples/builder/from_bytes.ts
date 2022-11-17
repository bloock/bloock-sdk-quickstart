import { RecordBuilder } from "@bloock/sdk"
import { Config } from "../../utils/config"
import { Sample } from "../../utils/sample"
import * as colors from 'colors'

Sample.run("builder_from_bytes", async (config: Config) => {
    let record = await RecordBuilder
        .fromBytes(Uint8Array.from([1, 2, 3, 4, 5]))
        .build();
    
    console.log(colors.green(`[✓] Record was created successfully`))

    let hash = await record.getHash()
    if (hash != "7d87c5ea75f7378bb701e404c50639161af3eff66293e9f375b5f17eb50476f4") {
        throw new Error("Unexpected hash received");
    }

    console.log(colors.green(`[✓] Hash: ${hash}`));
})