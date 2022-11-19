import com.bloock.sdk.Bloock;

public class FromLoader extends Sample 
{
    public void run(Config config) throws Exception
    {
        System.out.println( "FromLoader" );
    }

    FromLoader(String name) {
        super(name);
    }

    public static void main(String[] args) {
        new FromLoader("FromLoader");
    }
}
