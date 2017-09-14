install:
	go get ./cmd/...

lint:
	gometalinter --enable-all --line-length=120 --deadline 5m --exclude=data ./...

update-vendor:
	dep ensure
	dep ensure -update
	dep prune
