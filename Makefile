
GORUN := go run
GOBUILD := go build
BINARY_NAME := coolmate

DB_SOURCE := mysql://$(strip $(DB_USER)):$(strip $(DB_PASSWORD))@tcp($(strip $(DB_HOST)):$(strip $(DB_PORT)))/$(strip $(DB_NAME))?charset=utf8mb4&parseTime=True&loc=Local

format:
	@echo "Formatting code"
	@go fmt ./...
	@go vet ./...
	@echo "Formatted"

all: build

build: format
	@go build

run-member: format
	@if [ -z "$(ENV)" ]; then \
		echo "ENV variable not set. Defaulting to 'local'"; \
		export ENV=local; \
	fi
	@if [ -z "$(CONFIG_TYPE)" ]; then \
		echo "CONFIG_TYPE variable not set. Defaulting to 'env'"; \
		export CONFIG_TYPE=env; \
	fi
	@go run ./cmd/member_server/main.go

run-movie: format
	@if [ -z "$(ENV)" ]; then \
		echo "ENV variable not set. Defaulting to 'local'"; \
		export ENV=local; \
	fi
	@if [ -z "$(CONFIG_TYPE)" ]; then \
		echo "CONFIG_TYPE variable not set. Defaulting to 'env'"; \
		export CONFIG_TYPE=env; \
	fi
	@go run ./cmd/movie_server/main.go



start-db:
	@echo "Starting database..."
	@docker compose -f compose.db.yml up -d	

# example: make migrate-create name=add_users_table
migrate-create:
	@if [ -z "$(name)" ]; then \
		echo "Error: Migration name is required. Usage: make migrate-create name=<migration_name>"; \
		exit 1; \
	fi
	@echo "Creating migration file: $(name)"
	@migrate create -ext sql -dir migrations -seq $(name)

migrate-up:
	@echo "Applying migrations for ENV='$(ENV)'..."
	@echo "Using DB: mysql://$(DB_USER):******@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)"
	@migrate -database "$(DB_SOURCE)" -path migrations up

DB_COMPOSE := docker compose -f deployments/database/docker-compose.yml
API_COMPOSE := docker compose -f deployments/api/docker-compose.yml
CADDY_COMPOSE := docker compose -f deployments/caddy/docker-compose.yml

docker-build:
	@echo "Building API services..."
	@$(API_COMPOSE) build

docker-db-up:
	@echo "Starting database..."
	@$(DB_COMPOSE) up -d

docker-up:
	@echo "Starting database..."
	@$(DB_COMPOSE) up -d
	@echo "Starting API services..."
	@$(API_COMPOSE) up -d --build
	@echo "Starting Caddy reverse proxy..."
	@$(CADDY_COMPOSE) up -d

docker-down:
	@$(CADDY_COMPOSE) down
	@echo "Stopping API services..."
	@$(API_COMPOSE) down
	@echo "Stopping database..."
	@$(DB_COMPOSE) down

docker-clean: docker-down
	@docker volume rm coolmate-db_db-data || true

docker-logs-caddy:
	@$(CADDY_COMPOSE) logs -f

docker-logs-api:
	@$(API_COMPOSE) logs -f