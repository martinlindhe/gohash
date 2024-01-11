.PHONY: wasm

install:
	go install ./cmd/coder/...
	go install ./cmd/hasher/...
	go install ./cmd/findhash/...

bench:
	go test -bench=.

lint:
	golangci-lint run ./...

test-release:
	goreleaser --skip-publish --skip-validate --clean

test:
	go test -v

profile-mem:
	rm mem.prof
	go test --memprofile=mem.prof
	go tool pprof --alloc_space --text mem.prof
