import { RecordBuilder } from "@bloock/sdk"
import { Config } from "../../utils/config"
import { Sample } from "../../utils/sample"
import * as colors from 'colors'
import { readFileSync, writeFileSync } from "fs"

Sample.run("builder_from_file", async (config: Config) => {
    let file = readFileSync("../resources/dummy.pdf");
    let record = await RecordBuilder
        .fromFile(file)
        .build();

    console.log(colors.green(`[✓] Record was created successfully`))

    let hash = await record.getHash()
    if (hash != "43fa336d95e1634fee2d3031adc44ed9464472b94511af1facb87a1fee0b7e0e") {
        throw new Error("Unexpected hash received")
    }

    console.log(colors.green(`[✓] Hash: ${hash}`))

    // we can get the file back if needed
    writeFileSync("../resources/output.pdf", record.retrieve());
})