FROM docker.io/library/golang:alpine

WORKDIR /src/redis
COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /usr/bin/redis

CMD time redis