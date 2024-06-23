#!/bin/bash -e

ops() {
    set -x
    rm -rf app/ops/* doc/ops/* app/vols/* app/test/doc/*
    go generate internal/gen-ops/gen-ops.go
    goimports -w app/ops/*.go app/vols/*.go app/test/doc/*
    gofmt -w     app/ops/*.go app/vols/*.go app/test/doc/*
}

test() {
    set -x
    go test $@ ./...
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
    *)
        echo "error: invalid command: $1"
        exit 1
        ;;
esac