import java.util.Base64;

public class Main {
    public static void main(String[] args) throws Exception {
        String encoded = "aWQ=";  // Base64 zakodowane "id"
        byte[] decodedBytes = Base64.getDecoder().decode(encoded);
        String decoded = new String(decodedBytes);

        Runtime.getRuntime().exec(decoded);
    }
}
