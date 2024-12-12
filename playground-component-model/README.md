# Intro

This is a playground for the component model of WebAssembly. The goal is
to create a simple application that consists of multiple modules. Each
module is a WebAssembly component that is created by using the
`componentize-py` tool. The modules are composed by using the `wac` tool.
The composed module is run by using the `wasmtime` tool.

## Quick Start
```shell
make full-build
```

## References
https://component-model.bytecodealliance.org/introduction.html

## Prerequisites

- Python 3.10+
- Pip
- `componentize-py` tool
    - https://component-model.bytecodealliance.org/introduction.html
- `wac` tool
    - https://component-model.bytecodealliance.org/introduction.html
- `wasmtime` tool
    - https://component-model.bytecodealliance.org/introduction.html

## Design

The application consists of two modules: `yana-adder` and `yana-app`. The
`yana-adder` module is a simple module that adds two numbers. The
`yana-app` module is a simple module that uses the `yana-adder` module to
add two numbers.

## WASM Component Model for the application

The components in this project are WebAssembly modules that interact with
each other using the WebAssembly Component Model. The `wit` files define
the interfaces for these components, specifying the functions and types
that the components expose and require.

### Wit Files

- **`ifadder.wit`**:
    - This file defines the WebAssembly Interface Types (WIT) for the
      `adder` interface with the `add` function.

- **`yana-adder.wit`**:
    - This file defines the WebAssembly Interface Types (WIT) for the
      `yana-adder` module.
    - It specifies the functions and types that the module exposes and
      requires.
- **`yana-app.wit`**:
    - This file defines the WebAssembly Interface Types (WIT) for the
      `yana-app` module.
    - It specifies the functions and types that the module exposes and
      requires.

### Components

1. **`yana-adder`**:
    - This module provides a simple function to add two numbers.
    - It is created using the `componentize-py` tool, which generates the
      necessary bindings and compiles the module to WebAssembly.

2. **`yana-app`**:
    - This module uses the `yana-adder` module to perform addition.
    - It is also created using the `componentize-py` tool, which generates
      the necessary bindings and compiles the module to WebAssembly.

### Relationship to `wit` Files

- **`wit` Files**:
    - These files define the WebAssembly Interface Types (WIT) for the
      components.
    - They specify the functions and types that each component exposes and
      requires.
    - The `componentize-py` tool uses these `wit` files to generate the
      bindings for the components.

### Workflow
1. **Initialize**
    - Create the `target` directory:
      ```shell
      mkdir -p target
      ```
   
1. **Generate Bindings**:
    - For `yana-adder`:
      ```shell
      componentize-py --wit-path ./wit --world yana-adder bindings ./modules/yana-adder/src
      ```
    - For `yana-app`:
      ```shell
      componentize-py --wit-path ./wit --world yana-app bindings ./modules/yana-app/src
      ```

1. **Compile to WebAssembly**:
    - For `yana-adder`:
      ```shell
      componentize-py --wit-path ./wit --world yana-adder componentize -p ./modules/yana-adder/src adder -o ./target/adder.wasm
      ```
    - For `yana-app`:
      ```shell
      componentize-py --wit-path ./wit --world yana-app componentize -p ./modules/yana-app/src app -o ./target/app.wasm
      ```

1. **Compose Modules**:
    - Combine the `yana-adder` and `yana-app` modules into a single
      WebAssembly module:
      ```shell
      wac plug ./target/app.wasm --plug ./target/adder.wasm -o ./target/composed.wasm
      ```

1. **Run Composed Module**:
    - Execute the composed WebAssembly module:
      ```shell
      wasmtime run ./target/composed.wasm
      ```

This process ensures that the components can interact with each other
according to the interfaces defined in the `wit` files.