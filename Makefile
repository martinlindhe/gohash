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
