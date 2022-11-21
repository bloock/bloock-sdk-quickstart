package com.bloock.quickstart.samples.builder;

import com.bloock.quickstart.utils.Config;
import com.bloock.quickstart.utils.Logger;
import com.bloock.quickstart.utils.Sample;
import com.bloock.sdk.Bloock;

public class FromHex extends Sample 
{
    public void run(Config config) throws Exception
    {
        System.out.println( "FromHex" );
    }

    FromHex(String name) {
        super(name);
    }

    public static void main(String[] args) {
        new FromHex("FromHex");
    }
}
