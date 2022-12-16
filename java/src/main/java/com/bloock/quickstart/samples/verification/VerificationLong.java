package com.bloock.quickstart.samples.verification;

import com.bloock.quickstart.utils.Config;
import com.bloock.quickstart.utils.Logger;
import com.bloock.quickstart.utils.Sample;
import com.bloock.sdk.Bloock;
import com.bloock.sdk.builder.Builder;
import com.bloock.sdk.client.Client;
import com.bloock.sdk.entity.Anchor;
import com.bloock.sdk.entity.Network;
import com.bloock.sdk.entity.Proof;
import com.bloock.sdk.entity.Record;
import com.bloock.sdk.entity.RecordReceipt;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class VerificationLong extends Sample {
  public void run(Config config) throws Exception {
    Bloock.apiKey = config.apiKey;
    Client client = new Client();

    Record record = Builder.fromString("Hello world").build();

    ArrayList<Record> records = new ArrayList<>(Arrays.asList(record));

    List<RecordReceipt> receipts = client.sendRecords(records);

    // Once we sent a record, we can wait for it's anochor
    Logger.info("Waiting for anchor...");
    // we can optionally specify a timeout (if not set, default is 120000)
    Anchor anchor = client.waitAnchor(receipts.get(0).getAnchor(), 120000);
    Logger.success("Anchor " + anchor.getId() + " done!");

    // first we get the proof
    Proof proof = client.getProof(records);

    // then verify it
    String root = client.verifyProof(proof);

    // And finally validate the root
    // we can optionally specify a network (if not set, default is Ethereum Mainnet)
    long timestamp = client.validateRoot(root, Network.BLOOCK_CHAIN);

    Logger.success("Timestamp: " + timestamp);
  }

  VerificationLong(String name) {
    super(name);
  }

  public static void main(String[] args) {
    new VerificationLong("Verification long");
  }
}
