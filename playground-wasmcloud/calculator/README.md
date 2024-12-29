WasmCloud

This README explains how to build, deploy and execute the calculator application onto a local WASM Cloud installation. 

Run the following commands **inside the WASM toolchain shell.** For that, use Docker commands introduced in the README.md in the root directory of this repository.

0. When in the shell, make sure to navigate to the playground-wasmcloud/calculator directory.
```shell
cd playground-wasmcloud/calculator
```

1. Start the WASM Cloud environment
```shell
wash up -d
```
2. Build the calculator application
```shell
make build-all
```
3. Deploy the calculator application
```shell
make deploy-app
 ```
4. View the status of the application deployment. Wait until the status is "Deployed".
```shell
wash app list
```
5. Perform calculations by making HTTP calls using curl. For example, to add 20 and 5:
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

6. Delete the application (undeploy and remove from WASM Cloud)
```shell
make delete-app
```

7. Stop the WASM Cloud environment
```shell
wash down
```
