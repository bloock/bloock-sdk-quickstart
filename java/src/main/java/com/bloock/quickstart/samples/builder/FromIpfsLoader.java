package com.bloock.quickstart.samples.builder;

import com.bloock.quickstart.utils.Config;
import com.bloock.quickstart.utils.Logger;
import com.bloock.quickstart.utils.Sample;
import com.bloock.sdk.Bloock;
import com.bloock.sdk.builder.Builder;
import com.bloock.sdk.entity.IpfsLoader;
import com.bloock.sdk.entity.IpfsPublisher;
import com.bloock.sdk.entity.Record;

public class FromIpfsLoader extends Sample {
  public void run(Config config) throws Exception {
    Record record = Builder.fromString("Hello world").build();
    String hash = record.getHash();

    assert hash.equals("ed6c11b0b5b808960df26f5bfc471d04c1995b0ffd2055925ad1be28d6baadfd");

    Bloock.apiKey = config.apiKey;

    String result = record.publish(new IpfsPublisher());

    assert result.equals(hash);

    record = Builder.fromLoader(new IpfsLoader(result)).build();
    Logger.success("Record was created successfully");

    Logger.success("Hash: " + hash);
  }

  FromIpfsLoader(String name) {
    super(name);
  }

  public static void main(String[] args) {
    new FromIpfsLoader("From Ipfs Loader");
  }
}
