install:
	go get ./cmd/...

bench:
	go test -bench=.

lint:
	gometalinter --enable-all --line-length=180 --deadline 5m --exclude=vendor ./...

update-vendor:
	dep ensure
	dep ensure -update
	dep prune

release:
	goreleaser --rm-dist

test:
	go test -v

profile-mem:
	rm mem.prof
	go test --memprofile=mem.prof
	go tool pprof --alloc_space --text mem.prof
