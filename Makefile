PWD := $(shell pwd)
PKG := github.com/ckeyer/sloth
APP := sloth
DEV_IMAGE := ckeyer/dev

VERSION := $(shell cat VERSION.txt)
GIT_COMMIT := $(shell git rev-parse --short HEAD)
GIT_BRANCH := $(shell git rev-parse --abbrev-ref HEAD)

LD_FLAGS := -X $(PKG)/version.version=$(VERSION) -X $(PKG)/version.gitCommit=$(GIT_COMMIT) -w

local:
	go build -a -ldflags="$(LD_FLAGS)" -o bundles/$(APP) cli/main.go