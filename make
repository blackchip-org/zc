#!/bin/bash -e

GOFLAGS="-tags proj"

GEN_GO="app/ops/*.go app/vols/*.go app/test/docs/*.go app/test/ops/*.go"
GEN_MD="doc/ops/*.md"

function ops {
    rm -rf $GEN_GO $GEN_MD
    (
        set -x
        go generate internal/gen-ops/gen-ops.go
        go generate internal/gen-doc-tests/gen-doc-tests.go
    )
    goimports -w $GEN_GO
    gofmt     -w $GEN_GO

}

case "$1" in
    ops)
        ops
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
        set -x
        go test $GOFLAGS $@ ./...
        ;;
    bench)
        shift
        (ops)
        set -x
        go test $@ -benchmem -run=^$  -bench . github.com/blackchip-org/zc/v6/bench
        ;;
    wasm)
        (ops)
        set -x
        GOOS=js GOARCH=wasm go build -o web/zc.wasm cmd/wasm/main.go
        ;;
    emoji)
        set -x
        go generate internal/gen-emoji/gen-emoji.go
        ;;
    entity)
        set -x
        go generate internal/gen-entity/gen-entity.go
        ;;
    tz)
        set -x
        go generate internal/gen-tz/gen-tz.go
        ;;
    *)
        echo "error: invalid command: $1"
        exit 1
        ;;
esac