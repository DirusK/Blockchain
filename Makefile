.SILENT:
.EXPORT_ALL_VARIABLES:

BUILD_NAME = blockchain
APP_VERSION = v1.0.0

.PHONY: lint
lint:
	golangci-lint run -c golangci-lint.yml

.PHONY: test
test:
	go clean -testcache ./...
	go test ./...

.PHONY: check
check: test lint

.PHONY: build
build: test
	go build -ldflags \
		"-X 'main.version=$(APP_VERSION)'" \
		 -o ./cmd/$(BUILD_NAME) ./cmd/.

.PHONY: run
run:
	cd cmd; go run .

.PHONY: run-build
run-build: build
	cd cmd; ./$(BUILD_NAME)

.PHONY: clean
clean:
	go clean ./...
	rm -f ./cmd/$(BUILD_NAME)

.PHONY: rebuild
rebuild: clean build

.PHONY: mod
mod:
	go mod tidy
