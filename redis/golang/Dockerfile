FROM docker.io/library/golang:alpine

<<<<<<< HEAD
=======
WORKDIR /src/redis
>>>>>>> e372b74a2c0693509d98a69dc71e862278ba9a42
COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /usr/bin/redis

<<<<<<< HEAD
CMD time redis
=======
CMD time redis
>>>>>>> e372b74a2c0693509d98a69dc71e862278ba9a42
