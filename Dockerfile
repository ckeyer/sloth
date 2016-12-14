FROM alpine:edge

MAINTAINER Chuanjian Wang <me@ckeyer.com>

ENV GOPATH=/opt/gopath
ENV PATH=$PATH:$GOPATH/bin

RUN apk add --update git musl-dev go
RUN go get github.com/gordonklaus/ineffassign && \
	rm -rf $GOPATH/src/*

WORKDIR $GOPATH
EXPOSE 8080

