package com.bloock.quickstart.samples.node_proxy;

import com.bloock.sdk.Bloock;
import com.bloock.sdk.entity.Network;

import com.bloock.quickstart.utils.Config;
import com.bloock.quickstart.utils.Sample;

public class SetContractAddress extends Sample {
    public void run(Config config) throws Exception {
        // we can change the contract address by specifing the network and the new contract address
        Bloock.setContractAddress(Network.ETHEREUM_MAINNET, "some contract address");
    }

    SetContractAddress(String name) {
        super(name);
    }

    public static void main(String[] args) {
        new SetContractAddress("SetContractAddress");
    }
}
