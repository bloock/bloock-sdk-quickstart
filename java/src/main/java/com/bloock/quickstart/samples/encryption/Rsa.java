package com.bloock.quickstart.samples.encryption;

import com.bloock.quickstart.utils.Config;
import com.bloock.quickstart.utils.Logger;
import com.bloock.quickstart.utils.Sample;
import com.bloock.sdk.builder.Builder;
import com.bloock.sdk.client.Client;
import com.bloock.sdk.entity.KeyPair;
import com.bloock.sdk.entity.Record;
import com.bloock.sdk.entity.RsaDecrypter;
import com.bloock.sdk.entity.RsaEncrypter;

import java.nio.charset.StandardCharsets;

public class Rsa extends Sample {
  public void run(Config config) throws Exception {
    String payload = "This will be encrypted";
    Client sdk = new Client();

    KeyPair keypair = sdk.generateRsaKeyPair();

    Logger.info("The following payload will be encrypted: " + payload);

    // To encrypt we add an encrypter to the builder
    Record record = Builder.fromString(payload).withEncrypter(new RsaEncrypter(keypair.getPublicKey())).build();

    Logger.success("Encryption successful");

    Logger.success("Encrypted payload: " + new String(record.retrieve(), StandardCharsets.UTF_8));

    Logger.success("Record was created successfully");

    Logger.info("Trying to decrypt with the private key");

    // To decrypt we build a record from the encrypted record and add a decrypter
    Record decryptedRecord =
        Builder.fromRecord(record).withDecrypter(new RsaDecrypter(keypair.getPrivateKey())).build();

    Logger.success("Decryption successful");

    String hash = decryptedRecord.getHash();
    Logger.success("Hash: " + hash);

    Logger.success(
        "Decrypted payload: " + new String(decryptedRecord.retrieve(), StandardCharsets.UTF_8));

    assert decryptedRecord.retrieve().toString().equals(payload);

    Logger.info("Trying to decrypt with invalid key");

    Boolean exceptionWasThrown = false;
    try {
      Builder.fromRecord(record).withDecrypter(new RsaDecrypter("an invalid key")).build();
    } catch (Exception e) {
      exceptionWasThrown = true;
    }

    if (!exceptionWasThrown) {
      throw new Exception("The key was invalid so an error should've been returned!");
    }

    Logger.success("The key was invalid so the record could not be decrypted");
  }

  Rsa(String name) {
    super(name);
  }

  public static void main(String[] args) {
    new Rsa("RSA Encryption");
  }
}
