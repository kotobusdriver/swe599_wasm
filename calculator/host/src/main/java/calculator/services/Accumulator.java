package calculator.services;

public class Accumulator {
    private final ResultPrinter resultPrinter;
    private Integer result = 0;

    public Accumulator(ResultPrinter resultPrinter) {
        this.resultPrinter = resultPrinter;
    }

    public Integer read() {
        return result;
    }

    public void set(int newValue) {
        result = newValue;
        resultPrinter.print(result);
    }
}
