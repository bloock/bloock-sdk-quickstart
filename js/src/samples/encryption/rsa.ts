import { BloockClient, RecordBuilder, RsaDecrypter, RsaEncrypter } from "@bloock/sdk";
import { Config } from "../../utils/config";
import { Sample } from "../../utils/sample";
import { Logger } from "../../utils/logger";

Sample.run("RSA encryption", async (_: Config) => {
  let payload = "This will be encrypted";
  let sdk = new BloockClient();
  let keypair = await sdk.generateRsaKeyPair();

  Logger.warn(`The following payload will be encrypted: ${payload}`);

  // To encrypt we add an encrypter to the builder
  let encryptedRecord = await RecordBuilder.fromString(payload)
    .withEncrypter(new RsaEncrypter(keypair.publicKey))
    .build();

  Logger.success(`Encryption successfully`);

  let encryptedPayload = encryptedRecord.retrieve();
  Logger.success(
    `Encrypted payload: ${new TextDecoder().decode(encryptedPayload)}`
  );

  Logger.info("Trying to decrypt with the public key");

  // To decrypt we build a record from the encrypted record and add a decrypter
  let decryptedRecord = await RecordBuilder.fromRecord(encryptedRecord)
    .withDecrypter(new RsaDecrypter(keypair.privateKey))
    .build();

  Logger.success(`Decryption successfully`);

  let hash = await decryptedRecord.getHash();

  Logger.success(`Hash: ${hash}`);

  if (
    hash !== "6ede674d860c8593708d29f4484e87da36079f0bddf712e1e2bdc5ea0904f0ff"
  ) {
    throw new Error("Unexpected hash received");
  }

  let decryptedPayload = decryptedRecord.retrieve();
  Logger.success(
    `Decrypted payload: ${new TextDecoder().decode(decryptedPayload)}`
  );

  Logger.info("Trying to decrypt with invalid private key");

  let exceptionWasThrown = false;
  try {
    await RecordBuilder.fromRecord(encryptedRecord)
      .withDecrypter(new RsaDecrypter("an invalid key"))
      .build();
  } catch {
    exceptionWasThrown = true;
  }
  if (!exceptionWasThrown) {
    throw new Error(
      "The key was invalid so an error should've been returned!"
    );
  }

  Logger.success(
    `The key was invalid so the record could not be decrypted`
  );
});
