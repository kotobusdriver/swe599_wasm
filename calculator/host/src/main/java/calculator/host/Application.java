package calculator.host;

import calculator.services.*;
import org.graalvm.polyglot.Context;

public class Application {
    private final ResultPrinter resultPrinter;
    private final Accumulator accumulator;
    private final Adder adder;
    private final Subtractor subtractor;
    private final Multiplier multiplier;
    private final Divider divider;
    private final CommandInterpreter commandInterpreter;

    public Application(Context context) {
        resultPrinter = new ResultPrinter();
        accumulator = new Accumulator(resultPrinter);
        adder = new Adder(accumulator, context);
        subtractor = new Subtractor(accumulator, context);
        multiplier = new Multiplier(accumulator, context);
        divider = new Divider(accumulator, context);
        commandInterpreter = new CommandInterpreter(adder, subtractor, multiplier, divider, accumulator);
    }

    public void startCommandInterpreter() {
        commandInterpreter.start();
    }
}
