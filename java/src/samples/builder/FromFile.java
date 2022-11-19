import com.bloock.sdk.Bloock;

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
