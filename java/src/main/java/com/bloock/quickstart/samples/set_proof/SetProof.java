package com.bloock.quickstart.samples.set_proof;

import com.bloock.quickstart.utils.Config;
import com.bloock.quickstart.utils.Logger;
import com.bloock.quickstart.utils.Sample;
import com.bloock.sdk.builder.Builder;
import com.bloock.sdk.client.Client;
import com.bloock.sdk.entity.AnchorNetwork;
import com.bloock.sdk.entity.Proof;
import com.bloock.sdk.entity.ProofAnchor;
import com.bloock.sdk.entity.Record;
import java.util.Arrays;

public class SetProof extends Sample {
  public void run(Config config) throws Exception {
    Client sdk = new Client();

    Record record = Builder.fromString("Hello world").build();

    record.setProof(new Proof(
            Arrays.asList("ed6c11b0b5b808960df26f5bfc471d04c1995b0ffd2055925ad1be28d6baadfd"),
            Arrays.asList("ed6c11b0b5b808960df26f5bfc471d04c1995b0ffd2055925ad1be28d6baadfd"),
            "1010101",
            "0101010",
            new ProofAnchor(
                42L,
                Arrays.asList(
                    new AnchorNetwork(
                        "Ethereum",
                        "state",
                        "ed6c11b0b5b808960df26f5bfc471d04c1995b0ffd2055925ad1be28d6baadfd")),
                "ed6c11b0b5b808960df26f5bfc471d04c1995b0ffd2055925ad1be28d6baadfd",
                "succes")));

    Proof proof = sdk.getProof(Arrays.asList(record));
    Logger.success(
            "The following proof was set: " + 
            "leaves: " + proof.getLeaves() + ", " +
            "nodes: " + proof.getNodes() + ", " + 
            "depth: " + proof.getDepth() + ", " + 
            "bitmap: " + proof.getBitmap() + ", ");
  }

  SetProof(String name) {
    super(name);
  }

  public static void main(String[] args) {
    new SetProof("SetProof");
  }
}
