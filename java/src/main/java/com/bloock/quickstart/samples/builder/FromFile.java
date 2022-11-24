package com.bloock.quickstart.samples.builder;

import com.bloock.quickstart.utils.Config;
import com.bloock.quickstart.utils.Logger;
import com.bloock.quickstart.utils.Sample;
import com.bloock.sdk.builder.Builder;
import com.bloock.sdk.entity.Record;
import java.nio.file.Files;
import java.nio.file.Paths;

public class FromFile extends Sample {
  public void run(Config config) throws Exception {
    byte[] file = Files.readAllBytes(Paths.get("../resources/dummy.pdf"));

    Record record = Builder.fromFile(file).build();

    Logger.success("Record was created successfully");

    String hash = record.getHash();

    assert hash.equals("7d87c5ea75f7378bb701e404c50639161af3eff66293e9f375b5f17eb50476f4");

    Logger.success("Hash: " + hash);

    // we can get the file back if needed
    Files.write(Paths.get("../resources/output.pdf"), record.retrieve());

    Logger.success("File written");
  }

  FromFile(String name) {
    super(name);
  }

  public static void main(String[] args) {
    new FromFile("FromFile");
  }
}
