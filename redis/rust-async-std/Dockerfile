FROM docker.io/library/rust:alpine

COPY . .

RUN apk add --no-cache musl-dev && \
    cargo build --release

CMD time target/release/rust-async-std
