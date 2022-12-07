package com.bloock.quickstart.samples.node_proxy;

import com.bloock.quickstart.utils.Config;
import com.bloock.quickstart.utils.Sample;
import com.bloock.sdk.Bloock;
import com.bloock.sdk.entity.Network;

public class SetProvider extends Sample {
  public void run(Config config) throws Exception {
    // we can change the provider by specifing the network and the new url of the provider
    Bloock.setProvider(Network.ETHEREUM_MAINNET, "https://eth.someprovider.com");
  }

  SetProvider(String name) {
    super(name);
  }

  public static void main(String[] args) {
    new SetProvider("SetProvider");
  }
}
