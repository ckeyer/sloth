FROM alpine:edge

MAINTAINER Chuanjian Wang <me@ckeyer.com>

ENV GOPATH=/opt/gopath
ENV PATH=$PATH:$GOPATH/bin

RUN apk add --update git musl-dev go
# RUN go get golang.org/x/tools/... && \
# 	go get gopkg.in/alecthomas/gometalinter.v1 && \
# 	gometalinter.v1 --install --update && \
# 	rm -rf $GOPATH/src/*

WORKDIR $GOPATH
EXPOSE 8080

