package com.bloock.quickstart.samples.builder;

import com.bloock.quickstart.utils.Config;
import com.bloock.quickstart.utils.Sample;

public class FromJson extends Sample 
{
    public void run(Config config) throws Exception
    {
        System.out.println( "FromJson" );
    }

    FromJson(String name) {
        super(name);
    }

    public static void main(String[] args) {
        new FromJson("FromJson");
    }
}
