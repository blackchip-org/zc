.PHONY: all clean doc gen test

GOFLAGS=-tags proj

all: doc install

install: gen
	go install $(GOFLAGS) ./...

test: ops doc
	go test $(GOFLAGS) ./...

test-release: clean
	goreleaser release --skip-publish

release: clean gen
	goreleaser release

doc: index
	go generate internal/gen-doc/main.go

entity:
	go generate internal/gen-entity/main.go

index:
	go generate internal/gen-index/main.go

ops:
	go generate internal/gen-ops/main.go

tz:
	go generate internal/gen-tz/main.go

gen: ops doc index

clean:
	rm -rf dist
