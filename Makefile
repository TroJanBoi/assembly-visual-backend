# Simple Makefile for a Go project

# Build the application
all: 

# Run the application
run:
	@go run cmd/api/main.go

# Generate Swagger documentation
swagger:
	@echo "Generating Swagger documentation..."
	@swag init -g cmd/api/main.go -o cmd/api/docs
	@echo "Swagger documentation generated."

# Install Swagger CLI tool
swagger-install:
	@echo "Installing Swagger CLI tool..."
	@go install github.com/swaggo/swag/cmd/swag@latest
	@echo "Swagger CLI tool installed successfully!"