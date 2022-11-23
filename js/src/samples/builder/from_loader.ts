import { HostedLoader, HostedPublisher, RecordBuilder } from "@bloock/sdk";
import { Config } from "../../utils/config";
import { Logger } from "../../utils/logger";
import { Sample } from "../../utils/sample";

Sample.run("builder_from_loader", async (_: Config) => {
  let record = await RecordBuilder.fromString("Hello world").build();

  let hash = await record.getHash();
  if (
    hash !== "ed6c11b0b5b808960df26f5bfc471d04c1995b0ffd2055925ad1be28d6baadfd"
  ) {
    throw new Error("Unexpected hash received");
  }

  let result = await record.publish(new HostedPublisher());

  if (result !== hash) {
    throw new Error("Publish's result should be equal to the record's hash");
  }

  record = await RecordBuilder.fromLoader(new HostedLoader(result)).build();

  Logger.success(`Record was created successfully`);

  Logger.success(`Hash: ${hash}`);
});
