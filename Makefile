.PHONY: all clean doc gen test

GOFLAGS=-tags proj

all: doc install wasm

install: gen
	go install $(GOFLAGS) ./...

test: ops doc
	TZ=UTC go test -test.timeout=5s $(GOFLAGS) ./...

wasm:
	GOOS=js GOARCH=wasm go build -o web/zc.wasm cmd/wasm/main.go

serve:
	go run cmd/server/main.go

test-release: clean
	goreleaser release --skip-publish

release: clean gen
	goreleaser release

doc: index
	go generate internal/gen-doc/main.go

entity:
	go generate internal/gen-entity/main.go

emoji:
	go generate internal/gen-emoji/main.go

index:
	go generate internal/gen-index/main.go

ops:
	go generate internal/gen-ops/main.go

tz:
	go generate internal/gen-tz/main.go

gen: ops doc index

clean:
	rm -rf dist
