import { Bloock, BloockClient, Network, RecordBuilder } from "@bloock/sdk";
import { Config } from "../../utils/config";
import { Sample } from "../../utils/sample";
import { Logger } from "../../utils/logger";

Sample.run("Verification long", async (config: Config) => {
  // we set the API key and create a client
  Bloock.setApiKey(config.apiKey);
  let client = new BloockClient();

  let record = await RecordBuilder.fromString("Hello world").build();

  let hash = await record.getHash();
  let records = [hash];

  let receipts = await client.sendRecords(records);

  Logger.success(`Receipts: ${JSON.stringify(receipts)}`)
  // Once we sent a record, we can wait for it's anochor
  Logger.info(`Waiting for anchor...`);
  // we can optionally specify a timeout (if not set, default is 120000)
  let anchor = await client.waitAnchor(receipts[0].anchor, 120000);
  Logger.success(`Anchor ${anchor} done!`);

  // first we get the proof
  let proof = await client.getProof(records);

  // then verify it
  let root = await client.verifyProof(proof);

  // And finally validate the root
  // we can optionally specify a network (if not set, default is Ethereum Mainnet)
  let timestamp = await client.validateRoot(root, Network.BLOOCK_CHAIN);

  // We will recive a timestamp greater than 0 if the validation was successful
  Logger.success(`Timestamp: ${timestamp}`);
});
