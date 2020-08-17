FROM docker.io/crystallang/crystal:latest-alpine

COPY . .

RUN shards install --production --release && \
    shards build --production --release --progress --threads=$(nproc)

CMD time bin/crystal
