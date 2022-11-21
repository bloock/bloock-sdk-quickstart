package com.bloock.quickstart.samples.builder;

import com.bloock.quickstart.utils.Config;
import com.bloock.quickstart.utils.Logger;
import com.bloock.quickstart.utils.Sample;
import com.bloock.sdk.builder.Builder;
import com.bloock.sdk.entity.Record;

public class FromBytes extends Sample 
{
    public void run(Config config) throws Exception
    {
        Builder builder = Builder.fromString("hello world");
        Logger.info("builder" + ((builder != null) ? "true" : "false"));
        Record record = builder.build();
        Logger.info("record");
        String hash = record.getHash();
        Logger.info(hash);
        throw new Exception("FromBytes error");
    }

    FromBytes(String name) {
        super(name);
    }

    public static void main(String[] args) {
        new FromBytes("FromBytes");
    }
}
