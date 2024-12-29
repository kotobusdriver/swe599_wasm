export GOFLAGS=-buildvcs=false

.PHONY build-adder:
build-adder:
	cd modules/adder; \
	wash build;

.PHONY build-subtractor:
build-subtractor:
	cd modules/subtractor; \
	wash build;

.PHONY build-multiplier:
build-multiplier:
	cd modules/multiplier; \
	wash build;

.PHONY build-divider:
build-divider:
	cd modules/divider; \
	wash build;

.PHONY build-interpreter:
build-interpreter:
	cd modules/interpreter; \
	wit-deps; \
	wash build;

.PHONY build-all:
build-all: build-adder build-subtractor build-multiplier build-divider build-interpreter

.PHONY start-wasmcloud:
start-wasmcloud:
	wash up -d

.PHONY stop-wasmcloud:
stop-wasmcloud:
	wash down

.PHONY deploy-app:
deploy-app:
	wash app deploy --replace wadm.yaml

.PHONY delete-app:
delete-app:
	wash app delete wadm.yaml
