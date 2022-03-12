# syntax=docker/dockerfile:1

FROM golang:1.17.8-bullseye
ENV DEBIAN_FRONTEND noninteractive

# install build essentials
RUN apt-get update && \
  apt-get install -y wget build-essential pkg-config --no-install-recommends

# Install ImageMagick deps

RUN apt-get update && apt-get -q -y install libjpeg-dev libpng-dev libtiff-dev \
  libgif-dev libx11-dev --no-install-recommends

ENV IMAGEMAGICK_VERSION=7.1.0-27

RUN cd && \
  wget https://github.com/ImageMagick/ImageMagick/archive/${IMAGEMAGICK_VERSION}.tar.gz && \
  tar xvzf ${IMAGEMAGICK_VERSION}.tar.gz && \
  cd ImageMagick* && \
  ./configure \
  --without-magick-plus-plus \
  --without-perl \
  --disable-openmp \
  --with-gvc=no \
  --disable-docs && \
  make -j$(nproc) && make install && \
  ldconfig /usr/local/lib
RUN  rm -rf /var/lib/apt/lists/*

WORKDIR /app
ENV GO111MODULE=on
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go build -o /street ./cmd/
RUN chmod +x /street
EXPOSE 8088