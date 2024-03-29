package com.bloock.quickstart.utils;

public abstract class Sample {

  public static final String ANSI_RESET = "\u001B[0m";
  public static final String ANSI_YELLOW = "\u001B[33m";
  public static final String ANSI_RED = "\u001B[31m";
  public static final String ANSI_GREEN = "\u001B[32m";

  public Sample(String name) {
    Config config = new Config();
    config.apiKey = System.getenv("API_KEY");
    config.apiHost = System.getenv("API_HOST");

    Logger.info(name + ": Started");

    try {
      this.run(config);
      Logger.success(name + ": Successful");
    } catch (Exception e) {
      Logger.error(name + ": Failure");
      e.printStackTrace();
      Logger.error(name + ": " + e);
      System.exit(1);
    }
  }

  public abstract void run(Config config) throws Exception;
}
