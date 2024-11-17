# swe599_wasm
Polyglot distributed service-oriented architecture with web assembly 

November, 17

The initial commit contains a Java host application that performs basic arithmetic functions of addition, subtraction, multiplication, and division of two integers. The application contains seven sub-services: command interpreter, accumulator, results printer; as well as adder, subtractor, multiplier, and divider implemented as .wasm modules and called by the command interpreter. The command interpreter, accumulator and results printer are Java classes.
GraalWasm was used to embed .wasm files inside the Java host application. For this initial commit, hand written .wat files were used instead of compiling from other languages.
The application currently does not have safeguards (like division by 0). Float and long integers are not supported.

High-level diagram of the application
![image](https://github.com/user-attachments/assets/44f3f5bc-9cd9-4e7a-96b9-22170457d7d0)



Next steps: re-write arithmetic function modules in different programming languages, compile them to .wasm modules, and call from the host application.
