.PHONY: all doc test

all: doc install

install:
	go install $(GOFLAGS) ./...

doc:
	go generate internal/gen-index/main.go

test:
	go test $(GOFLAGS) ./...
