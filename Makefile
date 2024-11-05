BINARY_NAME=go-api
CMD_PATH=cmd/main.go

# Default target
.PHONY: all
all: build run

# Build the application
.PHONY: build
build:
	@ go build -o $(BINARY_NAME) $(CMD_PATH)

# Run the application
.PHONY: run
run:
	@ ./$(BINARY_NAME)

# Run the application in development mode
.PHONY: dev
dev:
	@ go run $(CMD_PATH)

# Clean the application
.PHONY: clean
clean:
	@ go clean
	@ rm -f $(BINARY_NAME)

# Regenerate Swagger documentation
.PHONY: swag
swag:
	@ swag init -g $(CMD_PATH)
