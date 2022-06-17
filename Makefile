# ==============================================================================
# DOCKER COMPOSE
# ==============================================================================
COMPOSE := docker-compose -f ./docker-compose.local.yml

load-env:
    export $(cut -d= -f1 .env)
docker-compose-up: load-env
	$(COMPOSE) up -d


# ==============================================================================
# SETUP
# ==============================================================================
install:
	echo Download go.mod dependencies
	go mod tidy
	go mod download

test-all:
	echo Run all test files in directories
	go run ./test/main.go

migrate:
	echo Migrate database schema
	go run cli/db_seed.go

swag:
	swag init -g ./cmd/serverd/main.go -o ./docs

# ==============================================================================
# RUN JOB
# ==============================================================================
run-api: install docker-compose-up migrate
	go run ./cmd/serverd/main.go

