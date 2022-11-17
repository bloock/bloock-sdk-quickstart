import { RecordBuilder } from "@bloock/sdk"
import { Config } from "../../utils/config"
import { Sample } from "../../utils/sample"
import * as colors from 'colors'

Sample.run("builder_from_json", async (config: Config) => {
    let record = await RecordBuilder
        .fromJson({ "hello" : "world" })
        .build();
    
    console.log(colors.green(`[✓] Record was created successfully`))

    let hash = await record.getHash()
    if (hash != "586e9b1e1681ba3ebad5ff5e6f673d3e3aa129fcdb76f92083dbc386cdde4312") {
        throw new Error("Unexpected hash received")
    }

    console.log(colors.green(`[✓] Hash: ${hash}`))
})