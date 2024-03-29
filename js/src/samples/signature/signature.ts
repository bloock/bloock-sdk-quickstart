import { Bloock, BloockClient, EcsdaSigner, RecordBuilder } from "@bloock/sdk";
import { Config } from "../../utils/config";
import { Logger } from "../../utils/logger";
import { Sample } from "../../utils/sample";

Sample.run("signature", async (config: Config) => {
  Bloock.setApiKey(config.apiKey);
  let client = new BloockClient();

  let keys = await client.generateKeys();

  let signedRecord = await RecordBuilder.fromString("Hello world")
    .withSigner(new EcsdaSigner(keys.privateKey))
    .build();

  Logger.success("Record was signed successfully");

  // we can add another signature with a different key
  keys = await client.generateKeys();

  Logger.info("Adding another signature with a common name");

  let name = "Some name";

  signedRecord = await RecordBuilder.fromRecord(signedRecord)
    .withSigner(new EcsdaSigner(keys.privateKey, { commonName: name }))
    .build();

  Logger.success("Record was signed successfully");

  Logger.success(`Hash: ${await signedRecord.getHash()}`);

  let signatures = await signedRecord.getSignatures();

  signatures.forEach((signature, i) => {
    Logger.success(`Signature ${i + 1}: ${signature.signature}`);
  });

  let retrievedName = await signatures[1].getCommonName();

  if (retrievedName != name) {
    throw new Error("The retrieved name is not the expected name")
  }

  Logger.success("Common name for signature is " + retrievedName)
});
