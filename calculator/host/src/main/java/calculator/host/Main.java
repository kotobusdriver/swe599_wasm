package calculator.host;

import org.graalvm.polyglot.Context;

public class Main {
    public static void main(String[] args) {
        try (Context context = Context.create()) {
            Application application = new Application(context);
            application.startCommandInterpreter();
        } catch (Throwable t) {
            System.out.println("Some error has occurred!");
            t.printStackTrace();
        }
    }
}
