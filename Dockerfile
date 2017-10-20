FROM alpine:edge

MAINTAINER Chuanjian Wang <me@ckeyer.com>

ENV GOPATH=/opt/gopath
ENV PATH=$PATH:$GOPATH/bin

RUN apk add --update git musl-dev go && \
	go get github.com/gordonklaus/ineffassign && \
	cp -a $GOPATH/bin/ /usr/local/bin/ && \
	rm -rf $GOPATH && \
	apk del go git

WORKDIR $GOPATH
EXPOSE 8080

