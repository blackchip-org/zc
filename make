#!/bin/bash -e

GOFLAGS="-tags proj"

GEN_GO="app/ops/*.go app/vols/*.go app/test/docs/*.go app/test/ops/*.go"
GEN_MD="doc/ops/*.md"

ops() {
    rm -rf $GEN_GO $GEN_MD
    go generate internal/gen-ops/gen-ops.go
    go generate internal/gen-doc-tests/gen-doc-tests.go
    goimports -w $GEN_GO
    gofmt     -w $GEN_GO
}

test() {
    set -x
    go test $GOFLAGS $@ ./...
}

bench() {
    set -x
    go test $@ -benchmem -run=^$  -bench . github.com/blackchip-org/zc/v6/bench
}

wasm() {
    set -x
    GOOS=js GOARCH=wasm go build -o web/zc.wasm cmd/wasm/main.go
}

case "$1" in
    ops)
        (ops)
        ;;
    run)
        shift
        go run cmd/zc/main.go $@
        ;;
    serve)
    	go run cmd/server/main.go
        ;;
    test)
        shift
        (ops)
        (test $@)
        ;;
    bench)
        shift
        (ops)
        (bench $@)
        ;;
    wasm)
        (ops)
        (wasm)
        ;;
    emoji)
        set -x
        go generate internal/gen-emoji/gen-emoji.go
        ;;
    *)
        echo "error: invalid command: $1"
        exit 1
        ;;
esac