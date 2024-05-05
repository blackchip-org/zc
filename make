#!/bin/bash -e

ops() {
    set -x
    go generate internal/gen-ops/gen-ops.go
    goimports -w ops/ops.go kinds/kinds.go volumes/*/{ops,volume}.go test/docs/*
    gofmt -w     ops/ops.go kinds/kinds.go volumes/*/{ops,volume}.go test/docs/*
}

test() {
    set -x
    go test $@ ./...
}

case "$1" in
    ops)
        (ops)
        ;;
    test)
        shift
        (test $@)
        ;;
    *)
        echo "error: invalid command: $1"
        exit 1
        ;;
esac