import com.bloock.sdk.Bloock;

public class FromString extends Sample 
{
    public void run(Config config) throws Exception
    {
        System.out.println( "FromString" );
    }

    FromString(String name) {
        super(name);
    }

    public static void main(String[] args) {
        new FromString("FromString");
    }
}
