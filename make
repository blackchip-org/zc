#!/bin/bash -e

ops() {
    set -x
    go generate internal/gen-ops/gen-ops.go
    goimports -w calc/ops/*.go calc/vols/*.go calc/test/docs/*
    gofmt -w     calc/ops/*.go calc/vols/*.go calc/test/docs/*
}

test() {
    set -x
    go test $@ ./...
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
    wasm)
        (ops)
        (wasm)
        ;;
    *)
        echo "error: invalid command: $1"
        exit 1
        ;;
esac