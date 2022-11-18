import { Bloock, BloockClient, RecordBuilder } from "@bloock/sdk"
import { Config } from "../../utils/config"
import { Sample } from "../../utils/sample"
import { Logger } from "../../utils/logger";

Sample.run("Send Data", async (config: Config) => {
    // we set the API key and create a client
    Bloock.setApiKey(config.apiKey);
    let client = new BloockClient();

    // we create an array of strings which will contain
    // the hashes of the records we want to send
    let records = []
    let record = await RecordBuilder
        .fromString("Hello world")
        .build();
        
    // we can get the hash of the record
    let hash = await record.getHash();

    // and append it to the array
    records.push(hash);

    // finally we can send the records
    let sendReceipts = await client.sendRecords(records)

    // we get a receipt with informationa about the transaction
    Logger.success(`Record receipts: ${sendReceipts}`);
})
