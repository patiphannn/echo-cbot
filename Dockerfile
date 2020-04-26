#!/bin/bash
FROM golang:1.14-alpine

LABEL maintainer="Patiphan Chaiya <patiphann@gmail.com>"

WORKDIR /app

COPY . /app

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./bin/echo-cbot .
CMD ["./bin/echo-cbot"]

EXPOSE 1323