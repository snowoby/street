# syntax=docker/dockerfile:1

FROM golang:1.17.6-alpine3.15
WORKDIR /app
ENV GO111MODULE=on
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go build -o /street ./cmd/street
EXPOSE 8088
CMD /street