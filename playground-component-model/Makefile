.PHONY init:
init:
	mkdir -p target

.PHONY clean:
clean:
	rm -rf ./target

.PHONY build-adder:
build-adder: init
	componentize-py --wit-path ./wit --world yana-adder componentize -p ./modules/yana-adder/src adder -o ./target/adder.wasm

.PHONY build-multiplier:
build-multiplier: init
	cd modules/yana-multiplier; \
	wkg wit build --wit-dir ../../wit; \
	go get go.bytecodealliance.org/cmd/wit-bindgen-go; \
	go run go.bytecodealliance.org/cmd/wit-bindgen-go generate -o . ./yana:calculator.wasm; \
	tinygo build -target=wasm-unknown -o ../../target/multiplier.wasm --wit-package ./yana/calculator/yana-multiplier/yana-multiplier.wit --wit-world yana-multiplier main.go;

.PHONY build-app:
build-app: init
	componentize-py --wit-path ./wit --world yana-app componentize -p ./modules/yana-app/src app -o ./target/app.wasm

.PHONY build-all:
build-all: build-adder build-app build-multiplier

.PHONY compose:
compose:
	wac plug ./target/app.wasm --plug ./target/adder.wasm -o ./target/app_add_int_1.wasm
	wac plug ./target/app_add_int_1.wasm --plug ./target/multiplier.wasm -o ./target/composed.wasm

.PHONY run:
run:
	wasmtime run ./target/composed.wasm

.PHONY full-build:
full-build: build-all compose run