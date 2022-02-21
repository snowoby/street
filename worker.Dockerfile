# syntax=docker/dockerfile:1

FROM golang:1.17.6-alpine3.15
WORKDIR /app
ENV GO111MODULE=on
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN apk --update add imagemagick imagemagick-dev build-base
RUN go build -o /tasker ./cmd/tasker
CMD /tasker