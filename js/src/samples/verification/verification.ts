import { Bloock, BloockClient, Network, RecordBuilder } from "@bloock/sdk";
import { Config } from "../../utils/config";
import { Sample } from "../../utils/sample";
import { Logger } from "../../utils/logger";

Sample.run("Verification", async (config: Config) => {
  // we set the API key and create a client
  Bloock.setApiKey(config.apiKey);
  let client = new BloockClient();

  let record = await RecordBuilder.fromString("Hello world").build();

  let records = [record];

  let receipts = await client.sendRecords(records);

  // Once we sent a record, we can wait for it's anochor
  Logger.info(`Waiting for anchor...`);
  // we can optionally specify a timeout (if not set, default is 120000)
  let anchor = await client.waitAnchor(receipts[0].anchor, 120000);
  Logger.success(`Anchor ${JSON.stringify(anchor)} done!`);

  // we can optionally specify a network (if not set, default is Ethereum Mainnet)
  let timestamp = await client.verifyRecords(records, Network.BLOOCK_CHAIN);
  Logger.success(`Timestamp: ${timestamp}`);
});
