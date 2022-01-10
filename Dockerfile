FROM golang:1.17 as local

WORKDIR /go/src/github.com/hiroyky/legato
COPY . .
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
RUN go get -u github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@v4.6.0
RUN go get -u github.com/volatiletech/sqlboiler/v4@v4.6.0
RUN go get -u github.com/99designs/gqlgen

FROM golang:1.17

WORKDIR /go/src/github.com/hiroyky/legato

COPY . .
RUN ln -s $(pwd) $GOPATH/src/
RUN go test ./...
RUN go build -o ./dist/app
RUN go build -o ./dist/import_sounds ./subsystem/import_sounds

EXPOSE 8080
CMD ["./dist/app"]
