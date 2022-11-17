import { RecordBuilder } from "@bloock/sdk"
import { Config } from "../../utils/config"
import { Sample } from "../../utils/sample"
import * as colors from 'colors'

Sample.run("builder_from_record", async (config: Config) => {
    let record = await RecordBuilder
        .fromString("Hello world")
        .build();

    record = await RecordBuilder
        .fromRecord(record)
        .build();

    console.log(colors.green(`[✓] Record was created successfully`));

    let hash = await record.getHash()
    if (hash != "ed6c11b0b5b808960df26f5bfc471d04c1995b0ffd2055925ad1be28d6baadfd") {
        throw new Error("Unexpected hash received");
    }

    console.log(colors.green(`[✓] Hash: ${hash}`));
})