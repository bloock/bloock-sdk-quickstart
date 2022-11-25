package com.bloock.quickstart.samples.verification;

import com.bloock.quickstart.utils.Config;
import com.bloock.quickstart.utils.Logger;
import com.bloock.quickstart.utils.Sample;
import com.bloock.sdk.Bloock;
import com.bloock.sdk.builder.Builder;
import com.bloock.sdk.client.Client;
import com.bloock.sdk.entity.Anchor;
import com.bloock.sdk.entity.Network;
import com.bloock.sdk.entity.Record;
import com.bloock.sdk.entity.RecordReceipt;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class Verification extends Sample {
  public void run(Config config) throws Exception {
    Bloock.apiKey = config.apiKey;
    Client client = new Client();

    Record record = Builder.fromString("Hello world").build();

    String hash = record.getHash();
    ArrayList<String> records = new ArrayList<>(Arrays.asList(hash));

    List<RecordReceipt> receipts = client.sendRecords(records);

    // Once we sent a record, we can wait for it's anochor
    Logger.info("Waiting for anchor...");
    // we can optionally specify a timeout (if not set, default is 120000)
    Anchor anchor = client.waitAnchor(receipts.get(0).getAnchor(), 120000);
    Logger.success("Anchor " + anchor.getId() + " done!");

    // we can optionally specify a network (if not set, default is Ethereum Mainnet)
    long timestamp = client.verifyRecords(records, Network.BLOOCK_CHAIN);
    Logger.success("Timestamp: " + timestamp);
  }

  Verification(String name) {
    super(name);
  }

  public static void main(String[] args) {
    new Verification("Verification");
  }
}
