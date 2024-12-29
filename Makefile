.PHONY run-toolchain:
run-toolchain:
	docker compose up --build -d

.PHONY start-shell:
start-shell:
	docker exec -it wasm_toolchain_container /bin/bash

.PHONY start-wasmcloud:
start-wasmcloud:
	docker compose -f ./playground-wasmcloud/local-env/docker-compose.yml up -d

.PHONY stop-wasmcloud:
stop-wasmcloud:
	docker compose -f ./playground-wasmcloud/local-env/docker-compose.yml down

.PHONY logs-wasmcloud:
logs-wasmcloud:
	docker logs -f local-env-wasmcloud-1