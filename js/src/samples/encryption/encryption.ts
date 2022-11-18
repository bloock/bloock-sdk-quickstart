import { AesDecrypter, AesEncrypter, RecordBuilder } from "@bloock/sdk"
import { Config } from "../../utils/config"
import { Sample } from "../../utils/sample"
import { Logger } from "../../utils/logger";

Sample.run("encryption", async (_: Config) => {
    let payload = "This wil be encrypted";
    let password = "a STRONG password";

    Logger.warn(`The following payload will be encrypted: ${payload}`);

    // To encrypt we add an encrypter to the builder
    let encryptedRecord = await RecordBuilder
        .fromString(payload)
        .withEncrypter(new AesEncrypter(password))
        .build();

    Logger.success(`Encryption successfully`);

    let encryptedPayload = encryptedRecord.retrieve();
    Logger.success(`Encrypted payload: ${new TextDecoder().decode(encryptedPayload)}`);

    Logger.info("Trying to decrypt with the valid password");

    // To decrypt we build a record from the encrypted record and add a decrypter
    let decryptedRecord = await RecordBuilder
        .fromRecord(encryptedRecord)
        .withDecrypter(new AesDecrypter(password))
        .build();

    Logger.success(`Decryption successfully`);

    let hash = await decryptedRecord.getHash();

    Logger.success(`Hash: ${hash}`);

    if (hash != "b6e6816e3c6180fcbda27048f033cf2b6f2a627864240c4c85558bcbece2a2e4") {
        throw new Error("Unexpected hash received");
    }

    let decryptedPayload = decryptedRecord.retrieve();
    Logger.success(`Decrypted payload: ${new TextDecoder().decode(decryptedPayload)}`);

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
