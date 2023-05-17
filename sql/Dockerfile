FROM postgres:15.1-alpine

LABEL author="ilker"
LABEL description="Postgres Image for Golang"
LABEL version="1.0"

COPY ./schema/*.sql /docker-entrypoint-initdb.d/
