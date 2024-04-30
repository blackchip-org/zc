#!/bin/bash

gen-ops() {
    set -x
    go generate internal/gen-ops/gen-ops.go
}

case "$1" in
    gen-ops) (gen-ops) ;;
esac