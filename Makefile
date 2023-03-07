.PHONY: doc test

BUILD_FLAGS = -tags=proj

doc:
	go generate internal/gen/index/main.go

test:
	go test $(BUILD_FLAGS) ./...
