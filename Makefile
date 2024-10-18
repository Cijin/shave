# Load environment variables from .env file
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

IMAGE_NAME=shave

# Simple helper to create secrets
# I always forget the command
new_secret:
	openssl rand -base64 64

init:
	sh init.sh

templ:
	@templ generate --watch

# Uses the rules in .air.toml file to run
server:
	@air

css:
	@npx tailwindcss -i ./public/css/input.css -o ./public/css/style.css --watch

dev:
	@make -j3 css templ server

test:
	go test ./...

update-mocks:
	@mockgen -source=./internal/database/querierExtended.go -destination=mocks/db/querierExtended.go 
	@mockgen -source=./pkg/aws/aws.go -destination=mocks/aws/aws.go

prod:
	@VERSION=$(shell git describe --tags --always --dirty) && \
	echo "Building Docker image with tag: $$VERSION" && \
	docker build -t $(IMAGE_NAME):$$VERSION .

db-up:
	@GOOSE_MIGRATION_DIR=$(MIGRATION_DIR) GOOSE_DBSTRING=$(SQL_URL) goose turso up

db-down:
	@GOOSE_MIGRATION_DIR=$(MIGRATION_DIR) GOOSE_DBSTRING=$(SQL_URL) goose turso down

db-down-all:
	@GOOSE_MIGRATION_DIR=$(MIGRATION_DIR) GOOSE_DBSTRING=$(SQL_URL) goose turso down-to 0
