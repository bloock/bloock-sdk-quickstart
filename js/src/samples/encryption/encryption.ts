import { AesDecrypter, AesEncrypter, RecordBuilder } from "@bloock/sdk"
import { Config } from "../../utils/config"
import { Sample } from "../../utils/sample"
import { Logger } from "../../utils/logger";

Sample.run("encryption", async (_: Config) => {
    let payload = "This wil be encrypted";
    let password = "a STRONG password";

    Logger.info(`The following payload will be encrypted: ${payload}`);

    // To encrypt we add an encrypter to the builder
    let encryptedRecord = await RecordBuilder
        .fromString(payload)
        .withEncrypter(new AesEncrypter(password))
        .build();

    Logger.success(`Encryption successfully`);
    Logger.success(`Encrypted payload: ${encryptedRecord.retrieve()}`);

    Logger.info("Trying to decrypt with the valid password");

    // To decrypt we build a record from the encrypted record and add a decrypter
    let decryptedRecord = await RecordBuilder
        .fromRecord(encryptedRecord)
        .withDecrypter(new AesDecrypter(password))
        .build();

    Logger.success(`Decryption successfully`);

    let hash = await decryptedRecord.getHash()
    if (hash != "7d87c5ea75f7378bb701e404c50639161af3eff66293e9f375b5f17eb50476f4") {
        throw new Error("Unexpected hash received");
    }

    Logger.success(`Hash: ${hash}`);
    Logger.success(`Decrypted payload: ${decryptedRecord.retrieve()}`);

    Logger.info("Trying to decrypt with invalid password");

    let exceptionWasThrown = false;
    try {
        await RecordBuilder
            .fromRecord(encryptedRecord)
            .withDecrypter(new AesDecrypter("an invalid password"))
            .build();
    } catch {
        exceptionWasThrown = true;
    }
    if (!exceptionWasThrown) {
        throw new Error("The password was invalid so an error should've been returned!");
    }

    Logger.success(`The password was invalid so the record could not be decrypted`);
})
