# syntax=docker/dockerfile:1

FROM golang:1.17.6-alpine3.15
WORKDIR /app
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories
RUN apk --update add imagemagick imagemagick-dev build-base

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go build -o /street ./cmd/street
RUN go build -o /worker ./cmd/worker
EXPOSE 8088
CMD /street & /worker