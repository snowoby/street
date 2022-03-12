# syntax=docker/dockerfile:1

FROM golang:1.17.8-bullseye
WORKDIR /app
ENV GO111MODULE=on
RUN apt-get update && apt-get install -y libmagickwand-dev && rm -rf /var/lib/apt/lists/*
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go build -o /street ./cmd/
RUN chmod +x /street
EXPOSE 8088