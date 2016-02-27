PKG := github.com/ckeyer/go-ci
APP := goci
GO :=  go 
VERSION := $(shell cat VERSION.txt)
LD_FLAGS := -X $(PKG)/version.version=$(VERSION)

default: build

build: 
	CGO_ENABLED=0 GOOS=linux GO15VENDOREXPERIMENT=1 $(GO) build -a -installsuffix nocgo -ldflags="$(LD_FLAGS)" -o bin/$(APP)

test: 
	 GO15VENDOREXPERIMENT=1 $(GO) test -ldflags="$(LD_FLAGS)" ./...
