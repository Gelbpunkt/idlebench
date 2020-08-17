#!/usr/bin/bash
podman pod create --name bench
podman run --rm -d --pod bench -v /home/jens/idlebench/schema.sql:/docker-entrypoint-initdb.d/schema.sql:z -e POSTGRES_PASSWORD="postgres" postgres:13-alpine
podman run --rm -d --pod bench redis:6-alpine
