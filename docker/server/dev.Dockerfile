FROM golang:1.13.3

RUN apt-get update && apt-get install -y postgresql-client wait-for-it

WORKDIR /go/src/github.com/jjjjackson/gqlgen-example

COPY . /go/src/github.com/jjjjackson/gqlgen-example

ENV GO111MODULE=on

CMD [ "go", "run", "main.go" ]
