import { Bloock, BloockClient, Network, RecordBuilder } from "@bloock/sdk"
import { Config } from "../../utils/config"
import { Sample } from "../../utils/sample"
import * as colors from 'colors'

Sample.run("Verification", async (config: Config) => {
    // we set the API key and create a client
    Bloock.setApiKey(process.env["API_KEY"]);
    let client = new BloockClient();

    let record = await RecordBuilder
        .fromString("Hello world")
        .build();

    let hash = await record.getHash();
    let records = [hash];

    let receipts = await client.sendRecords(records);

    // Once we sent a record, we can wait for it's anochor
    console.log(colors.yellow(`[+] Waiting for anchor...`));
    // we can optionally specify a timeout (if not set, default is 120000) 
    let anchor = await client.waitAnchor(receipts[0].anchor, 120000);
    console.log(colors.green(`[✓] Anchor ${anchor} done!`));

    // we can optionally specify a network (if not set, default is Ethereum Mainnet)
    let timestamp = await client.verifyRecords(records, Network.ETHEREUM_MAINNET);
    console.log(timestamp);

    // we get a receipt with informationa about the transaction
    console.log(colors.green(`[✓] Record receipts: ${sendReceipts}`));
})

Sample.run("Verification long", async (config: Config) => {
    // we set the API key and create a client
    Bloock.setApiKey(process.env["API_KEY"]);
    let client = new BloockClient();

    let record = await RecordBuilder
        .fromString("Hello world")
        .build();

    let hash = await record.getHash();
    let records = [hash];

    let receipts = await client.sendRecords(records);

    // Once we sent a record, we can wait for it's anochor
    console.log(colors.yellow(`[+] Waiting for anchor...`));
    // we can optionally specify a timeout (if not set, default is 120000) 
    let anchor = await client.waitAnchor(receipts[0].anchor, 120000);
    console.log(colors.green(`[✓] Anchor ${anchor} done!`));

    // first we get the proof
    let proof = await client.getProof(records);

    // then verify it
    let root = await client.verifyProof(proof);

    // And finally validate the root
    // we can optionally specify a network (if not set, default is Ethereum Mainnet)
    let timestamp = await client.validateRoot(root, Network.ETHEREUM_MAINNET)

    // We will recive a timestamp greater than 0 if the validation was successful
    console.log(colors.green(`[✓] Timestamp: ${timestamp}`));
})