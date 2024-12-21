# swe599_wasm
Polyglot distributed service-oriented architecture with web assembly 

November, 17

The initial commit contains a Java host application that performs basic arithmetic functions of addition, subtraction, multiplication, and division of two integers. The application contains seven sub-services: command interpreter, accumulator, results printer; as well as adder, subtractor, multiplier, and divider implemented as .wasm modules and called by the command interpreter. The command interpreter, accumulator and results printer are Java classes.
GraalWasm was used to embed .wasm files inside the Java host application. For this initial commit, hand written .wat files were used instead of compiling from other languages.
The application currently does not have safeguards (like division by 0). Float and long integers are not supported.

High-level diagram of the application
![image](https://github.com/user-attachments/assets/44f3f5bc-9cd9-4e7a-96b9-22170457d7d0)



Next steps: re-write arithmetic function modules in different programming languages, compile them to .wasm modules, and call from the host application.


November, 24

Three of the four arithmetic function modules are re-written in C, Rust, and TypeScript and compiled to .wasm files
with their respective compilers - Emscripten, Cargo, and AssemblyScript. 
Against initial expectation, it was difficult to compile Python and Go codes into standalone wasm modules. When compilation
finally did succeed, it was not posible to run the resulting wasm modules in GraalWasm because some runtime
dependencies could not be resolved.
This means that the tooling for compilation from Go to Wasm produced .wasm components with Go-specific dependencies, though 
it would most likely run in a Go environment. The same problem was for the Pythone-to-wasm generated modules.

As the next step, I will investigate using the Wasm Component Model tool chain - https://component-model.bytecodealliance.org/-  
to create independent wasm components that will (hopefully) integrate nicer with the wasm ecosystem.

December, 12

I am exploring the WebAssembly component model. The current goal is to build a simple application consisting of two WebAssembly modules:
yana-adder: A module that adds two numbers.
yana-app: A module that uses yana-adder to perform addition.

The modules interact via interfaces defined in wit files using the WebAssembly Interface Types (WIT). These interfaces specify the functions and data types exposed and required by each module.
The workflow is as follows:
Define interfaces: create wit files for both modules.
Generate bindings: using componentize-py, generate bindings for Python and compile the modules to WebAssembly.
Compose modules: combine the two modules into one using the WebAssembly Composition (WAC)  tool.
Run the appl: in wasmtime environment, execute the composed WebAssembly modules.

December 15

Start the toolchain container.
```shell
docker-compose up --build
```
Start a shell in the container.
```shell
docker exec -it wasm_toolchain_container /bin/bash
```