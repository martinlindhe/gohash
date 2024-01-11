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
