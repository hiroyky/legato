# Makefile

LOCAL_DOCKER_COMPOSE_FILE=./docker-compose.local.yaml
PROD_DOCKER_COMPOSE_FILE=./docker-compose.prod.yaml

init:
	docker compose --file $(LOCAL_DOCKER_COMPOSE_FILE) build --no-cache
	docker compose --file $(LOCAL_DOCKER_COMPOSE_FILE) up -d

restart:
	docker compose --file $(LOCAL_DOCKER_COMPOSE_FILE) stop && docker compose --file $(LOCAL_DOCKER_COMPOSE_FILE)  up -d

gen: gen_sqlboil gen_gql

gen_gql:
	docker compose --file $(LOCAL_DOCKER_COMPOSE_FILE) exec legato go run github.com/99designs/gqlgen

gen_sqlboil:
	docker compose --file $(LOCAL_DOCKER_COMPOSE_FILE) exec legato go run github.com/volatiletech/sqlboiler/v4 mysql

bash:
	docker compose --file $(LOCAL_DOCKER_COMPOSE_FILE) exec legato bash

fmt:
	docker compose --file $(LOCAL_DOCKER_COMPOSE_FILE) exec legato go fmt ./...

test:
	docker compose --file $(LOCAL_DOCKER_COMPOSE_FILE) exec legato go test ./...

clean:
	rm -rf $(DST_DIR) ./dist

build_prod:
	docker compose --file $(PROD_DOCKER_COMPOSE_FILE) build --no-cache
	docker compose --file $(PROD_DOCKER_COMPOSE_FILE) up -d

build: _build_app _build_import_sounds

_build_app:
	go build -o ./dist/app

_build_import_sounds:
	go build -o ./dist/import_sounds ./subsystem/import_sounds
