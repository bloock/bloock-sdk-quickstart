import { RecordBuilder } from "@bloock/sdk"
import { Config } from "../../utils/config"
import { Sample } from "../../utils/sample"
import { readFileSync, writeFileSync } from "fs"
import { Logger } from "../../utils/logger"

Sample.run("builder_from_file", async (_: Config) => {
    let file = readFileSync("../resources/dummy.pdf");
    let record = await RecordBuilder
        .fromFile(file)
        .build();

    Logger.success(`Record was created successfully`)

    let hash = await record.getHash()
    if (hash != "43fa336d95e1634fee2d3031adc44ed9464472b94511af1facb87a1fee0b7e0e") {
        throw new Error("Unexpected hash received")
    }

    Logger.success(`Hash: ${hash}`);

    // we can get the file back if needed
    writeFileSync("../resources/output.pdf", record.retrieve());
    Logger.success(`File written`);
})
