import { AnchorNetwork, BloockClient, Proof, ProofAnchor, RecordBuilder } from "@bloock/sdk";
import { Config } from "../../utils/config";
import { Sample } from "../../utils/sample";
import { Logger } from "../../utils/logger";

Sample.run("Send Data", async (_: Config) => {
  let record = await RecordBuilder.fromString("Hello world").build();

  await record.setProof(new Proof(
    ["ed6c11b0b5b808960df26f5bfc471d04c1995b0ffd2055925ad1be28d6baadfd"],
    ["ed6c11b0b5b808960df26f5bfc471d04c1995b0ffd2055925ad1be28d6baadfd"],
    "1010101",
    "0101010",
    new ProofAnchor(
      42,
      [
        new AnchorNetwork(
          "Ethereum",
          "state",
          "ed6c11b0b5b808960df26f5bfc471d04c1995b0ffd2055925ad1be28d6baadfd"
        )
      ],
      "ed6c11b0b5b808960df26f5bfc471d04c1995b0ffd2055925ad1be28d6baadfd",
      "success"
    )
  ));

  let proof = await new BloockClient().getProof([record]);

  Logger.success(`The following proof was set: ${JSON.stringify(proof)}`);
});
