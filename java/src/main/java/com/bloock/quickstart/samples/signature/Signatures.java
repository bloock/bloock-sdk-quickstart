package com.bloock.quickstart.samples.signature;

import java.util.List;

import com.bloock.quickstart.utils.Config;
import com.bloock.quickstart.utils.Logger;
import com.bloock.quickstart.utils.Sample;
import com.bloock.sdk.Bloock;
import com.bloock.sdk.builder.Builder;
import com.bloock.sdk.client.Client;
import com.bloock.sdk.entity.EcsdaSigner;
import com.bloock.sdk.entity.Keys;
import com.bloock.sdk.entity.Record;
import com.bloock.sdk.entity.Signature;
import com.bloock.sdk.entity.SignerArgs;

public class Signatures extends Sample {
  public void run(Config config) throws Exception {
    Bloock.apiKey = config.apiKey;
    Client client = new Client();

    Keys keys = client.generateKeys();

    Record signedRecord =
        Builder.fromString("Hello world").withSigner(new EcsdaSigner(keys.getPrivateKey())).build();

    Logger.success("Record was signed successfully");

    // we can add another signature with a different key
    Logger.info("Adding another signature with common name");
    keys = client.generateKeys();

    String name = "Some name";

    signedRecord = Builder.fromRecord(signedRecord)
        .withSigner(new EcsdaSigner(new SignerArgs(keys.getPrivateKey(), name)))
        .build();

    Logger.success("Record was signed successfully");

    Logger.success("Hash: " + signedRecord.getHash());

    List<Signature> signatures = signedRecord.getSignatures();
    signatures.stream()
        .forEach(
            signature -> {
              Logger.success("Signature " + signature.getSignature());
            });

    String retrievedName = signatures.get(1).getCommonName();
    assert retrievedName.equals(name);

    Logger.success("Common name for signature is " + retrievedName);
  }

  Signatures(String name) {
    super(name);
  }

  public static void main(String[] args) {
    new Signatures("Signature");
  }
}
