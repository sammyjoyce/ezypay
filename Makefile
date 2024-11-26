# {{project_name}} Makefile
# Author: {{author_name}} <{{author_email}}>

.PHONY: proto clean dev dev-api dev-web build test watch

# Load environment variables from .env file
-include .env

# Default values if not set in .env
PORT ?= 8080
GRPC_PORT ?= 50051
GRPC_HOST ?= 0.0.0.0

proto: install-proto-deps
	@chmod +x scripts/generate_go_protos.sh scripts/generate_web_protos.sh
	@./scripts/generate_go_protos.sh
	@./scripts/generate_web_protos.sh

install-proto-deps:
	@echo "Installing protobuf dependencies..."
	@command -v protoc >/dev/null 2>&1 || (echo "Installing protoc..." && brew install protobuf)
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@cd web && pnpm install

clean:
	rm -rf internal/gen/* web/app/api/generated/protos/*

build:
	@echo "Building..."
	@go build -o main ./cmd/api

test:
	@echo "Running tests..."
	@go test ./... -v

watch:
	@if command -v air > /dev/null; then \
		air -c configs/.air.toml; \
	else \
		read -p "Go's 'air' is not installed. Install it? [Y/n] " choice; \
		if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
			go install github.com/air-verse/air@latest; \
			air -c configs/.air.toml; \
		else \
			echo "You chose not to install air. Exiting..."; \
			exit 1; \
		fi; \
	fi

# Development targets
dev: proto
	@echo "Starting development environment..."
	@make -j2 dev-api dev-web

dev-api: build
	@echo "Starting API server on port $(PORT)..."
	@./main

dev-web:
	@echo "Starting web server..."
	@cd web && pnpm dev