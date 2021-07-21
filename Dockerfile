FROM golang:1.16

WORKDIR /go/src/github.com/hiroyky/legato

COPY . .
RUN ln -s $(pwd) $GOPATH/src/
RUN make test build

EXPOSE 8080
CMD ["./dist/app"]
