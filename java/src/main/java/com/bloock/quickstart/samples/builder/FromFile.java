package com.bloock.quickstart.samples.builder;

import com.bloock.quickstart.utils.Config;
import com.bloock.quickstart.utils.Sample;

public class FromFile extends Sample 
{
    public void run(Config config) throws Exception
    {
        System.out.println( "FromFile" );
    }

    FromFile(String name) {
        super(name);
    }

    public static void main(String[] args) {
        new FromFile("FromFile");
    }
}