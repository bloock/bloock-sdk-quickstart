import com.bloock.sdk.Bloock;

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
