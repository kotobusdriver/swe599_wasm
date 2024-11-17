package calculator.services;

import java.util.Arrays;
import java.util.Optional;
import java.util.Scanner;

public class CommandInterpreter {
    private final Adder adder;
    private final Subtractor subtractor;
    private final Multiplier multiplier;
    private final Divider divider;
    private final Accumulator accumulator;

    public CommandInterpreter(Adder adder, Subtractor subtractor, Multiplier multiplier, Divider divider, Accumulator accumulator) {
        this.adder = adder;
        this.subtractor = subtractor;
        this.multiplier = multiplier;
        this.divider = divider;
        this.accumulator = accumulator;
    }

    public void start() {
        Scanner scanner = new Scanner(System.in);
        System.out.println("Enter commands in the format: command [arg1] [arg2]");
        System.out.println("Type 'exit' to quit.");

        while (true) {
            System.out.print("> ");
            String input = scanner.nextLine().trim();

            if (input.equalsIgnoreCase("exit")) {
                break;
            }

            // Parse input
            String[] parts = input.split("\\s+", 3); // Split into at most 3 parts

            Optional<Command> command;
            if (parts.length > 0) {
                command = Command.resolve(parts[0]);
                if (command.isEmpty()) {
                    continue;
                }
            } else {
                continue;
            }

            int arg1 = 0;
            int arg2 = 0;

            if (parts.length == 2) {
                arg1 = accumulator.read();
                arg2 = Integer.parseInt(parts[1]);
            } else if (parts.length == 3) {
                arg1 = Integer.parseInt(parts[1]);
                arg2 = Integer.parseInt(parts[2]);
            }

            process(command.get(), arg1, arg2);

        }

        scanner.close();
    }

    enum Command {
        RESET("reset"),
        ADD("add"),
        SUB("sub"),
        MUL("mul"),
        DIV("div"),
        EXIT("exit");

        public final String cmd;

        Command(String cmd) {
            this.cmd = cmd;
        }

        public static Optional<Command> resolve(String text) {
            return Arrays.stream(values()).filter(c -> c.cmd.equals(text)).findAny();
        }
    }

    private void process(Command command, Integer arg1, Integer arg2) {
        switch(command) {
            case RESET -> accumulator.set(0);
            case EXIT -> System.exit(0);
            case ADD -> adder.process(arg1, arg2);
            case SUB -> subtractor.process(arg1, arg2);
            case MUL -> multiplier.process(arg1, arg2);
            case DIV -> divider.process(arg1, arg2);
        }
    }
}
