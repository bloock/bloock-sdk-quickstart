package com.bloock.quickstart.samples.builder;

import com.bloock.quickstart.utils.Config;
import com.bloock.quickstart.utils.Logger;
import com.bloock.quickstart.utils.Sample;
import com.bloock.sdk.builder.Builder;
import com.bloock.sdk.entity.Record;

public class FromBytes extends Sample {
  public void run(Config config) throws Exception {
    Record record = Builder.fromBytes(new byte[] {1, 2, 3, 4, 5}).build();

    Logger.success("Record was created successfully");

    String hash = record.getHash();

    assert hash.equals("7d87c5ea75f7378bb701e404c50639161af3eff66293e9f375b5f17eb50476f4");

    Logger.success("Hash: " + hash);
  }

  FromBytes(String name) {
    super(name);
  }

  public static void main(String[] args) {
    new FromBytes("FromBytes");
  }
}
