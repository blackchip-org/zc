.PHONY: all clean doc gen test

all: doc install

install:
	go install $(GOFLAGS) ./...

doc:
	go generate internal/gen-doc/main.go
	go generate internal/gen-index/main.go

test: gen
	go test $(GOFLAGS) ./...

test-release: clean
	goreleaser release --skip-publish

release: clean gen
	goreleaser release

gen-ops:
	go generate internal/gen-ops/main.go

gen: gen-ops doc
	go generate internal/gen-tz/main.go

clean:
	rm -rf dist
