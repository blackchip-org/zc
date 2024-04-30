#!/bin/bash

gen-ops() {
    set -x
    go generate internal/gen-ops/gen-ops.go
}

test() {
    set -x
    go test ./...
}

case "$1" in
    gen-ops)
        (gen-ops)
        ;;
    test)
        (test)
        ;;
    *)
        echo "error: invalid command: $1"
        exit 1
        ;;
esac