import { AesDecrypter, AesEncrypter, RecordBuilder } from "@bloock/sdk"
import { Config } from "../../utils/config"
import { Sample } from "../../utils/sample"
import * as colors from 'colors'

Sample.run("encryption", async (config: Config) => {
    let payload = "This wil be encrypted";
    let password = "a STRONG password";

    console.log(colors.yellow(`[+] The following payload will be encrypted: ${payload}`));

    let encryptedRecord = await RecordBuilder
        .fromString(payload)
        .withEncrypter(new AesEncrypter(password))
        .build();

    console.log(colors.green(`[✓] Encryption successfully`));
    console.log(colors.green(`[✓] Encrypted payload: ${encryptedRecord.retrieve()}`));

    console.log(colors.yellow("[+] Trying to decrypt with the valid password"));

    let decryptedRecord = await RecordBuilder
        .fromRecord(encryptedRecord)
        .withDecrypter(new AesDecrypter(password))
        .build();

    console.log(colors.green(`[✓] Decryption successfully`));

    let hash = await decryptedRecord.getHash()
    if (hash != "7d87c5ea75f7378bb701e404c50639161af3eff66293e9f375b5f17eb50476f4") {
        throw new Error("Unexpected hash received");
    }

    console.log(colors.green(`[✓] Hash: ${hash}`));
    console.log(colors.green(`[✓] Decrypted payload: ${decryptedRecord.retrieve()}`));

    console.log(colors.yellow("[+] Trying to decrypt with invalid password"));

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
    
    console.log(colors.green(`[✓] The password was invalid so the record could not be decrypted`));

})
