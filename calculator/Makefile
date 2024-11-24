EM_SDK_PATH:=~/temp/wasm/emsdk/emsdk_env.sh

.PHONY build-adder:
build-adder:
	source ${EM_SDK_PATH}
	emcc --no-entry services/adder/adder.c -o services/adder/adder.wasm


.PHONY build-subtractor:
build-subtractor:
	cargo build  --manifest-path=services/subtractor/Cargo.toml --target wasm32-wasi --release


.PHONY build-multiplier:
build-multiplier:
	asc services/multiplier/multiplier.ts -o services/multiplier/multiplier.wasm --optimize

.PHONY build-divider:
build-divider:
	wat2wasm divider.wat