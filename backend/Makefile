APP_NAME = url-shorty
CMD_PATH = ./cmd/server

.PHONY: all build run test clean lint docker

all: build

## Build the Go binary
backend:
	@echo "Building $(APP_NAME)..."
	go build -o bin/$(APP_NAME) $(CMD_PATH)

## Run the app (rebuild if needed)
run-backend: build
	@echo "Running $(APP_NAME)..."
	./bin/$(APP_NAME)

## Run tests
test:
	@echo "Running tests..."
	go test ./... -v

## Format and lint
lint:
	@echo "Linting..."
	gofmt -s -w .
	go vet ./...

## Clean build files
clean:
	@echo "Cleaning..."
	rm -rf bin/
