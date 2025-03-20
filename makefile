.PHONY: default run build test docs clean

# Variables
APP_NAME=gopportunities

# Tasks
default: run-with-docs

run:
	@go run main.go
run-with-docs: docs run
build:
	@go build -o $(APP_NAME) main.go
test:
	@go test ./...
docs:
	@swag init
clean:
	@rm -rf $(APP_NAME)
	@rm -rf ./docs