.PHONY: all clean doc gen test

all: doc install

install:
	go install $(GOFLAGS) ./...

doc:
	go generate internal/gen-index/main.go

test:
	go test $(GOFLAGS) ./...

test-release: clean
	goreleaser release --skip-publish

gen:
	go generate internal/gen-tz/main.go

clean:
	rm -rf dist
