package com.bloock.quickstart.samples.send_data;

import com.bloock.quickstart.utils.Config;
import com.bloock.quickstart.utils.Logger;
import com.bloock.quickstart.utils.Sample;
import com.bloock.sdk.Bloock;
import com.bloock.sdk.builder.Builder;
import com.bloock.sdk.client.Client;
import com.bloock.sdk.entity.Record;
import com.bloock.sdk.entity.RecordReceipt;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class SendData extends Sample {
  public void run(Config config) throws Exception {
    Bloock.apiKey = config.apiKey;
    Client client = new Client();

    // we create an ArrayList of strings which will contain
    // the hashes of the records we want to send
    Record record = Builder.fromString("Hello world").build();
    String hash = record.getHash();
    ArrayList<String> records = new ArrayList<>(Arrays.asList(hash));

    // finally we can send the records
    List<RecordReceipt> receipts = client.sendRecords(records);

    // we get a receipt with informationa about the transaction
    Logger.success("Record receipts: ");
    receipts.stream()
        .forEach(
            x -> {
              Logger.success(
                  "anchor: "
                      + x.getAnchor()
                      + ", client: "
                      + x.getClient()
                      + ", record: "
                      + x.getRecord()
                      + ", status: "
                      + x.getStatus());
            });
  }

  SendData(String name) {
    super(name);
  }

  public static void main(String[] args) {
    new SendData("SendData");
  }
}
