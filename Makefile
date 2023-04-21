.PHONY: all doc gen test

all: doc install

install:
	go install $(GOFLAGS) ./...

doc:
	go generate internal/gen-index/main.go

test:
	go test $(GOFLAGS) ./...

gen:
	go generate internal/gen-tz/main.go
