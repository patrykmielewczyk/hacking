public class Main {
    public static void main(String[] args) throws Exception {
        char[] cmd = {105, 100};  // ASCII for "id"
        String command = new String(cmd);

        Runtime.getRuntime().exec(command);
    }
}
