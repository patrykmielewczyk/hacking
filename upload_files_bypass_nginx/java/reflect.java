import java.lang.reflect.Method;

public class Main {
    public static void main(String[] args) throws Exception {
        Method method = Runtime.class.getMethod("exec", String.class);
        method.invoke(Runtime.getRuntime(), "id");
    }
}
