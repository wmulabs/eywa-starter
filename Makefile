.PHONY: up down run chat tidy build test

COMPOSE = docker compose -f deployments/docker-compose.yaml

# Send a sample message to the running agent. Override: make chat MSG="your text".
MSG ?= what analytics tools are in the catalog?

# Start MongoDB + Redis in the background.
up:
	$(COMPOSE) up -d

# Stop the infrastructure.
down:
	$(COMPOSE) down

# Run the API locally (loads .env). Requires `make up` first.
run:
	@set -a; . ./.env; set +a; go run ./cmd/api

# Send a sample message to the running agent (see MSG above).
chat:
	@curl -s localhost:8080/api/v1/events/chat \
		-H 'Content-Type: application/json' \
		-d '{"user":"u1","message":"$(MSG)"}' | jq

# Sync dependencies.
tidy:
	go mod tidy

# Compile the API binary.
build:
	go build -o bin/api ./cmd/api

test:
	go test ./... -race
