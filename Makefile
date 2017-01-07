PWD := $(shell pwd)
PKG := github.com/ckeyer/sloth
APP := sloth

DEV_IMAGE := ckeyer/dev
DEV_UI_IMAGE := ckeyer/dev:node

VERSION := $(shell cat VERSION.txt)
GIT_COMMIT := $(shell git rev-parse --short HEAD)
GIT_BRANCH := $(shell git rev-parse --abbrev-ref HEAD)

LD_FLAGS := -X $(PKG)/version.version=$(VERSION) -X $(PKG)/version.gitCommit=$(GIT_COMMIT) -w

NET := $(shell docker network inspect cknet > /dev/zero && echo "--net cknet --ip 172.16.1.8" || echo "")
UI_NET := $(shell docker network inspect cknet > /dev/zero && echo "--net cknet --ip 172.16.1.9" || echo "")

HASH := $(shell which sha1sum || which shasum)

# 连接url ： [mongodb:// ][user:pass@]host1[:port1][,host2[:port2],...][/database][?options]
MGO_URL := mongodb://ckeyer:X4etb83XtjlXz@u3.mj:27017/sloth

init:
	which gometalinter || (go get github.com/alecthomas/gometalinter && gometalinter -i)

local:
	go build -a -ldflags="$(LD_FLAGS)" -o bundles/$(APP) cli/main.go
	$(HASH) bundles/$(APP)

test:
	go test -ldflags="$(LD_FLAGS)" $$(go list ./... |grep -v "vendor")

run: local
	MGO_URL=$(MGO_URL) \
	DEBUG=true \
	UI_DIR="../sloth-ui/dist" \
	bundles/$(APP) eval ./

dev:
	docker run --rm -it \
	 --name $(APP)-dev \
	 -p 8000:8000 \
	 $(NET) \
	 -v /var/run/docker.sock:/var/run/docker.sock \
	 -v $(PWD)/..:/opt/gopath/src/$(PKG)/.. \
	 -w /opt/gopath/src/$(PKG) \
	 $(DEV_IMAGE) bash

dev-ui:
	docker run --rm -it \
	 --name $(APP)-ui-dev \
	 -p 8080:8080 \
	 -v $(PWD)/ui:/opt/$(APP) \
	 -w /opt/$(APP) \
	 $(DEV_UI_IMAGE) bash

reg:
	curl -i "http://localhost:8000/api/signup" -d '{"name":"ckeyer", "password":"wangcj", "email":"dev@ckeyer.com"}'

login:
	curl -i "http://localhost:8000/api/login" -d '{"password":"wangcj", "email":"dev@ckeyer.com"}'