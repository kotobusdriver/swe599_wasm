package calculator.services;

import calculator.host.Main;
import calculator.host.Operation;
import org.graalvm.polyglot.Context;
import org.graalvm.polyglot.Source;
import org.graalvm.polyglot.Value;

import java.io.IOException;
import java.net.URL;

public class Adder implements Operation {
    private final Accumulator accumulator;
    private final Value wasmCall;
    public Adder(Accumulator accumulator, Context context) {
        this.accumulator = accumulator;
        this.wasmCall = setupWasm(context);
    }

    private Value setupWasm(Context context) {
        try {
            URL wasmFile = Main.class.getResource("/wasm/adder.wasm");
            String moduleName = "adder";
            context.eval(Source.newBuilder("wasm", wasmFile).name(moduleName).build());
            return context.getBindings("wasm").getMember(moduleName).getMember("process");
        } catch (IOException e) {
            throw new RuntimeException(e);
        }
    }

    @Override
    public void process(Integer arg1, Integer arg2) {
        int result = wasmCall.execute(arg1, arg2).asInt();
        accumulator.set(result);
    }
}
