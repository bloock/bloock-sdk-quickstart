import { RecordBuilder } from "@bloock/sdk"
import { Config } from "../../utils/config"
import { Sample } from "../../utils/sample"
import { Logger } from "../../utils/logger"

Sample.run("builder_from_hex", async (_: Config) => {
    let record = await RecordBuilder
        .fromHex("1234567890abcdef")
        .build();
    
    Logger.success(`Record was created successfully`)

    let hash = await record.getHash()
    if (hash != "ed8ab4fde4c4e2749641d9d89de3d920f9845e086abd71e6921319f41f0e784f") {
        throw new Error("Unexpected hash received")
    }

    Logger.success(`Hash: ${hash}`);
})
