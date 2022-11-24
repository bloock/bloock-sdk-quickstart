package com.bloock.quickstart.samples.builder;

import com.bloock.quickstart.utils.Config;
import com.bloock.quickstart.utils.Logger;
import com.bloock.quickstart.utils.Sample;
import com.bloock.sdk.builder.Builder;
import com.bloock.sdk.entity.Record;

public class FromJson extends Sample {
  public void run(Config config) throws Exception {
    Record record = Builder.fromJson("{\"hello\":\"world\"}").build();

    Logger.success("Record was created successfully");

    String hash = record.getHash();

    assert hash.equals("586e9b1e1681ba3ebad5ff5e6f673d3e3aa129fcdb76f92083dbc386cdde4312");

    Logger.success("Hash: " + hash);
  }

  FromJson(String name) {
    super(name);
  }

  public static void main(String[] args) {
    new FromJson("FromJson");
  }
}
