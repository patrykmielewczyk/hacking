import java.lang.reflect.Method;

public class Main {
    public static void main(String[] args) throws Exception {
        String funcName = "getRuntime";
        Method method = Runtime.class.getMethod(funcName);
        Runtime runtime = (Runtime) method.invoke(null);
        runtime.exec("id");
    }
}
