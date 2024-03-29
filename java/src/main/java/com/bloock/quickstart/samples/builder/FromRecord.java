package com.bloock.quickstart.samples.builder;

import com.bloock.quickstart.utils.Config;
import com.bloock.quickstart.utils.Logger;
import com.bloock.quickstart.utils.Sample;
import com.bloock.sdk.builder.Builder;
import com.bloock.sdk.entity.Record;

public class FromRecord extends Sample {
  public void run(Config config) throws Exception {
    Record record = Builder.fromString("Hello world").build();

    record = Builder.fromRecord(record).build();

    Logger.success("Record was created successfully");

    String hash = record.getHash();

    assert hash.equals("ed6c11b0b5b808960df26f5bfc471d04c1995b0ffd2055925ad1be28d6baadfd");

    Logger.success("Hash: " + hash);
  }

  FromRecord(String name) {
    super(name);
  }

  public static void main(String[] args) {
    new FromRecord("FromRecord");
  }
}
