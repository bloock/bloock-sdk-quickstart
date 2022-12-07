import { Bloock, Network } from "@bloock/sdk";
import { Config } from "../../utils/config";
import { Sample } from "../../utils/sample";

Sample.run("set provider", async (_: Config) => {
  // we can change the provider by specifing the network and the new url of the provider
  Bloock.setProvider(Network.ETHEREUM_MAINNET, "https://eth.someprovider.com");
});

Sample.run("set contract address", async (_: Config) => {
  // we can change the contract address by specifing the network and the new contract address
  Bloock.setContractAddress(Network.ETHEREUM_MAINNET, "some contract address");
});
