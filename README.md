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

The WASM development requires a set of toolchains to be present on the developer's machine. To help with the installation steps, 
here is a Docker environment that contains all the necessary tools. The remaining part of the document assumes the use of this  WASM toolchain shell.


December 29
![image](https://github.com/user-attachments/assets/89580d23-dc48-4bd0-852f-f8771688b1b3)

Here are the complete instructions on how to run the developed application in WASM Cloud.

1. Start the toolchain container. Be patient, the first time it will take a while to build the image.
```shell
docker-compose up --build -d
```
Start a shell in the container.
```shell
docker exec -it wasm_toolchain_container /bin/bash
```

2. When in the shell, make sure to navigate to the playground-wasmcloud/calculator directory.
```shell
cd playground-wasmcloud/calculator
```

3. Start the WASM Cloud environment
```shell
wash up -d
```
4. Build the calculator application
```shell
make build-all
```
5. Deploy the calculator application
```shell
make deploy-app
 ```
6. View the status of the application deployment. Wait until the status is "Deployed".
```shell
wash app list
```
7. Perform calculations by making HTTP calls using curl. For example, to add 20 and 5:
```shell
curl "http://localhost:8000?op=add&a=20&b=5"
```
```
Result = 25
```

For example, to multiply 20 and 5:
```shell
curl "http://localhost:8000?op=mul&a=20&b=5"
```
```
Result = 100
```
Allowed operations are add, sub, mul, and div. The operands a and b are unsigned integers.

8. Delete the application (undeploy and remove from WASM Cloud)
```shell
make delete-app
```

9. Stop the WASM Cloud environment
```shell
wash down
```

