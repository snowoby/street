# syntax=docker/dockerfile:1

FROM golang:1.17.6-alpine3.15
WORKDIR /app
ENV GO111MODULE=on
RUN apk --no-cache --update add imagemagick imagemagick-dev build-base && rm -rf /var/cache/apk/*
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go build -o /street ./cmd/
RUN chmod +x /street
EXPOSE 8088