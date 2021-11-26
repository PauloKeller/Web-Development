FROM golang:alpine

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download 

ADD . .

CMD CGO_ENABLED=0 go test ./...

