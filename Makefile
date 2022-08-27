.DEFAULT_GOAL := help

DOCKER_TAG := latest

.PHONY: build-production
build-prod: ## Build docker image to deploy
	docker build -t aselab/ase-lab-backend:${DOCKER_TAG} \
		--target deploy ./

.PHONY: build
build: ## Build docker image to local development
	docker compose build --no-cache

.PHONY: up-prod
up-prod: ## Do docker compose up with production
	docker compose -f docker-compose.prod.yml up -d

.PHONY: up
up: ## Do docker compose up with hot reload
	docker compose up -d

.PHONY: down
down: ## Do docker compose down
	docker compose down

.PHONY: logs
logs: ## Tail docker compose logs
	docker compose logs -f

.PHONY: ps
ps: ## Check container status
	docker compose ps

.PHONY: test
test: ## Execute tests
	go test -race -shuffle=on ./...

.PHONY: gen
gen:
	go generate ./...

.PHONY: help
help: ## Show options
	@cat $(MAKEFILE_LIST) | grep -E '^[a-zA-Z_-]+:.*?## .*$$' | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

