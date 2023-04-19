.PHONY: wasm

install:
	go install ./cmd/coder/...
	go install ./cmd/hasher/...
	go install ./cmd/findhash/...

bench:
	go test -bench=.

lint:
	golangci-lint run ./...

release:
	goreleaser --rm-dist

test:
	go test -v

profile-mem:
	rm mem.prof
	go test --memprofile=mem.prof
	go tool pprof --alloc_space --text mem.prof

wasm:
	# tinygo 0.11 fails to compile math/big
	# curl -L https://raw.githubusercontent.com/tinygo-org/tinygo/master/targets/wasm_exec.js -o wasm/coder/wasm_exec.js
	tinygo build -o wasm/coder/wasm.wasm -target=wasm wasm/coder/wasm.go
