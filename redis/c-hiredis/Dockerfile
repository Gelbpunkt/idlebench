FROM docker.io/library/alpine:edge

RUN apk add --no-cache gcc hiredis-dev json-c-dev musl-dev

COPY . .

RUN gcc main.c -o /usr/bin/main -lhiredis -ljson-c

CMD time main
