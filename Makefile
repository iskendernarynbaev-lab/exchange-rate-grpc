APP_NAME=app
GO=go
MIGRATIONS_DIR=internal/storage/migrations
GOOSE=$(GO) run github.com/pressly/goose/v3/cmd/goose@v3.24.1

ifneq (,$(wildcard .env))
include .env
export
endif

.PHONY: build test docker-build run lint migrate-up migrate-down migrate-status migrate-create

build:
	$(GO) build -o $(APP_NAME) ./cmd/app

test:
	$(GO) test ./...

docker-build:
	docker build -t grinex-rates:latest .

run:
	$(GO) run ./cmd/app

lint:
	docker run --rm -v $(PWD):/app -w /app golangci/golangci-lint:v2.1.6 golangci-lint run ./...

migrate-up:
	$(GOOSE) -dir $(MIGRATIONS_DIR) postgres "$(DATABASE_URL)" up

migrate-down:
	$(GOOSE) -dir $(MIGRATIONS_DIR) postgres "$(DATABASE_URL)" down

migrate-status:
	$(GOOSE) -dir $(MIGRATIONS_DIR) postgres "$(DATABASE_URL)" status

migrate-create:
	@if [ -z "$(NAME)" ]; then echo "NAME is required. Example: make migrate-create NAME=add_users_table"; exit 1; fi
	$(GOOSE) -dir $(MIGRATIONS_DIR) create $(NAME) sql
