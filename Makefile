###########################################
.EXPORT_ALL_VARIABLES:
VERSION_PKG := github.com/ergoapi/util/version
ROOT_DIR := $(CURDIR)
BUILD_DIR := $(ROOT_DIR)/_output
BIN_DIR := $(BUILD_DIR)/bin
GO111MODULE = on
GOPROXY = https://goproxy.cn,direct
GOSUMDB = sum.golang.google.cn
IMAGE           ?= ghcr.io/ysicing/bigcat

BUILD_RELEASE   ?= $(shell cat version.txt || echo "0.0.1")
BUILD_DATE := $(shell date "+%Y%m%d")
GIT_COMMIT := $(shell git rev-parse --short HEAD || echo "abcdefgh")
APP_VERSION := ${BUILD_RELEASE}-${BUILD_DATE}-${GIT_COMMIT}

LDFLAGS := "-w \
	-X $(VERSION_PKG).release=$(APP_VERSION) \
	-X $(VERSION_PKG).gitVersion=$(APP_VERSION) \
	-X $(VERSION_PKG).gitCommit=$(GIT_COMMIT) \
	-X $(VERSION_PKG).gitBranch=$(GIT_BRANCH) \
	-X $(VERSION_PKG).buildDate=$(BUILD_DATE) \
	-X $(VERSION_PKG).gitTreeState=core \
	-X $(VERSION_PKG).gitMajor=0 \
	-X $(VERSION_PKG).gitMinor=0"

GO_BUILD_FLAGS+=-ldflags $(LDFLAGS)
GO_BUILD := go build $(GO_BUILD_FLAGS)

##########################################################################

help: ## this help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

gencopyright: ## add copyright
	@bash hack/build/gencopyright.sh

fmt: ## fmt code
	gofmt -s -w .
	goimports -w .
	@echo gofmt -l
	@OUTPUT=`gofmt -l . 2>&1`; \
	if [ "$$OUTPUT" ]; then \
		echo "gofmt must be run on the following files:"; \
        echo "$$OUTPUT"; \
        exit 1; \
    fi

lint: ## lint code
	@echo golangci-lint run --skip-files \".*test.go\" -v ./...
	@OUTPUT=`command -v golangci-lint >/dev/null 2>&1 && golangci-lint run --skip-files ".*test.go"  -v ./... 2>&1`; \
	if [ "$$OUTPUT" ]; then \
		echo "go lint errors:"; \
		echo "$$OUTPUT"; \
	fi

doc: ## doc
	hack/build/gendocs.sh

default: gencopyright doc fmt lint ## fmt code

static: ## 构建ui
	hack/build/ui.sh ${APP_VERSION}

build: ## build binary
	@echo "build bin ${GIT_VERSION} $(GIT_COMMIT) $(GIT_BRANCH) $(BUILD_DATE) $(GIT_TREE_STATE)"
	$(GO_BUILD) -o $(BIN_DIR)/bigcat

docker: static ## 构建镜像
	docker build -t ${IMAGE}:${BUILD_RELEASE} .
	docker tag ${IMAGE}:${BUILD_RELEASE} ${IMAGE}
	docker push ${IMAGE}
	docker push ${IMAGE}:${BUILD_RELEASE}

clean: ## clean
	rm -rf $(BIN_DIR)

.PHONY : build clean
