# Makefile

gengql:
	go run github.com/99designs/gqlgen

build:
	go build -o app

fmt:
	go fmt ./...
