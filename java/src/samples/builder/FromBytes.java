import com.bloock.sdk.Bloock;
import com.bloock.sdk.builder.Builder;
import com.bloock.sdk.client.Client;
import com.bloock.sdk.entity.Record;

public class FromBytes extends Sample 
{
    public void run(Config config) throws Exception
    {
        Bloock.apiKey = "";
        Client client = new Client();

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
