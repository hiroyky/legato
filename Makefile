# Makefile

gengql:
	go run github.com/99designs/gqlgen

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
