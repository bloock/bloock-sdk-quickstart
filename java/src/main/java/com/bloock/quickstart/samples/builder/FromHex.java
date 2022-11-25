package com.bloock.quickstart.samples.builder;

import com.bloock.quickstart.utils.Config;
import com.bloock.quickstart.utils.Logger;
import com.bloock.quickstart.utils.Sample;
import com.bloock.sdk.builder.Builder;
import com.bloock.sdk.entity.Record;

public class FromHex extends Sample {
  public void run(Config config) throws Exception {
    Record record = Builder.fromHex("1234567890abcdef").build();

    Logger.success("Record was created successfully");

    String hash = record.getHash();

    assert hash.equals("ed8ab4fde4c4e2749641d9d89de3d920f9845e086abd71e6921319f41f0e784f");

    Logger.success("Hash: " + hash);
  }

  FromHex(String name) {
    super(name);
  }

  public static void main(String[] args) {
    new FromHex("FromHex");
  }
}
