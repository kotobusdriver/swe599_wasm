.PHONY run-toolchain:
run-toolchain:
	docker compose up --build -d

.PHONY start-shell:
start-shell:
	docker exec -it wasm_toolchain_container /bin/bash