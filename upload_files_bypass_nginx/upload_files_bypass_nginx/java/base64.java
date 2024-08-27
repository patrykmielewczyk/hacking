import java.util.Base64;

public class Main {
    public static void main(String[] args) throws Exception {
        String encoded = "aWQ=";  // base64 coded "id"
        byte[] decodedBytes = Base64.getDecoder().decode(encoded);
        String decoded = new String(decodedBytes);

        Runtime.getRuntime().exec(decoded);
    }
}
