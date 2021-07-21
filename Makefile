# Makefile

prepare: _prepare_sqlboiler_mysql

_prepare_sqlboiler_mysql:
	go get github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql

generate: gen_sqlboil gen_gql

gen_gql:
	go run github.com/99designs/gqlgen

gen_sqlboil:
	go run github.com/volatiletech/sqlboiler/v4 mysql

build:
	go build -o ./dist/app

fmt:
	go fmt ./...

dev:
	docker-compose stop
	docker-compose build
	docker-compose up -d
	docker-compose ps

test:
	go test ./...

clean:
	rm -rf $(DST_DIR) ./dist
