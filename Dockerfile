FROM golang:1.7-alpine
RUN set -ex && \
    apk add --no-cache git build-base

WORKDIR /usr/local/go/src/github.com/evrone/docker-machine-vscale

ADD . /usr/local/go/src/github.com/evrone/docker-machine-vscale
RUN make fetch
