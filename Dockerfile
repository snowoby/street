# syntax=docker/dockerfile:1

FROM golang:1.17.6-alpine3.15
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
ENV GO111MODULE=on \
        GOPROXY=https://goproxy.io,direct
RUN go mod download
COPY . .
RUN go build -o /street ./cmd/street
EXPOSE 8088
CMD /street