package com.bloock.quickstart.samples.builder;

import com.bloock.quickstart.utils.Config;
import com.bloock.quickstart.utils.Sample;

public class FromRecord extends Sample 
{
    public void run(Config config) throws Exception
    {
        System.out.println( "FromRecord" );
    }

    FromRecord(String name) {
        super(name);
    }

    public static void main(String[] args) {
        new FromRecord("FromRecord");
    }
}
