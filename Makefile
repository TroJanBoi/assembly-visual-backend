# ---------------------------------------
# Makefile for Go Project with Colors 🎨
# ---------------------------------------

# ANSI colors
GREEN  := \033[1;32m
BLUE   := \033[1;34m
YELLOW := \033[1;33m
RED    := \033[1;31m
RESET  := \033[0m

.PHONY: all build run swagger swagger-install clean test

# ---------------------------------------
# Default Target: Build Everything
# ---------------------------------------
all: test build

# ---------------------------------------
# Build the Application
# ---------------------------------------
build:
	@echo "$(YELLOW)🔨 Building the application...$(RESET)"
	@go build -o main cmd/api/main.go
	@echo "$(GREEN)✅ Build complete!$(RESET)"

# ---------------------------------------
# Run the Application
# ---------------------------------------
run:
	@echo "$(YELLOW)🚀 Starting application...$(RESET)"
	@go run cmd/api/main.go

# ---------------------------------------
# Test the Application
# ---------------------------------------
test:
	@echo "$(YELLOW)🧪 Running tests...$(RESET)"
	@go clean -testcache
	@go test ./test/... -v
	@echo "$(GREEN)✅ Tests complete!$(RESET)"

# ---------------------------------------
# Generate Swagger Documentation
# ---------------------------------------
swagger:
	@echo "$(BLUE)📄 Generating Swagger documentation...$(RESET)"
	@swag init -g cmd/api/main.go -o cmd/api/docs
	@echo "$(GREEN)✅ Swagger docs generated at cmd/api/docs$(RESET)"

# ---------------------------------------
# Install Swagger CLI Tool
# ---------------------------------------
swagger-install:
	@echo "$(BLUE)⬇️ Installing Swagger CLI tool...$(RESET)"
	@go install github.com/swaggo/swag/cmd/swag@latest
	@echo "$(GREEN)✅ Swagger CLI tool installed!$(RESET)"

# ---------------------------------------
# Clean Up Build Artifacts
# ---------------------------------------
clean:
	@echo "$(RED)🧹 Cleaning build artifacts...$(RESET)"
	@rm -f main
	@echo "$(GREEN)✅ Clean complete!$(RESET)"

