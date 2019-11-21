SHELL := /bin/bash

# 当前目录
CUR_DIR := $(shell pwd)
# Go 模块名称
GO_MODULE_NAME := github.com/coffee377/voc-admin
# 构建命令
BUILD := go build
# 构建名称
BUILD_NAME := api
# 构建输出
OUT = -o bin/$(BUILD_NAME)_$(GOOS)_$(GOARCH)$(EXT)
# https://studygolang.com/articles/9463
FLAGS = -ldflags "-w \
	-X '$(VERSION_PKG).GitTag=$(GIT_TAG)' \
	-X '$(VERSION_PKG).BuildTime=$(BUILD_TIME)' \
	-X '$(VERSION_PKG).GitCommit=$(GIT_COMMIT)' \
	-X '$(VERSION_PKG).GitTreeState=$(GIT_TREE_STATE)' \
	-X '$(VERSION_PKG).Os=$(GOOS)' \
	-X '$(VERSION_PKG).Arch=$(GOARCH)'"
# Go 程序入口
PKG := ./cmd
BUILD_COMMAND = $(BUILD) $(OUT) ${FLAGS} $(PKG)

# 版本工具所在包
VERSION_PKG := $(GO_MODULE_NAME)/pkg/version
BUILD_TIME := $(shell date "+%F %T")
GIT_TAG := $(shell if [ "`git describe --tags --abbrev=0 2>/dev/null`" != "" ];then git describe --tags --abbrev=0; \
	else git log --date=format:'%Y-%m-%d %H:%M:%S' --pretty=format:'%cd' -n 1 ; fi)
GIT_COMMIT := $(shell git log --pretty=format:'%H' -n 1)
GIT_TREE_STATE := $(shell if git status| grep -q 'clean';then echo clean; else echo dirty; fi)

# 构建版本目标
BUILD_TARGET = windows_386 windows_amd64 linux_386 \
	linux_amd64 linux_arm linux_arm64 \
	darwin_386 darwin_amd64 darwin_arm darwin_arm64

# 系统类型变量设置
windows% : GOOS = windows
linux% : GOOS = linux
darwin% : GOOS = darwin

# 处理器变量设置
%386: GOARCH = 386
%arm: GOARCH = arm
%arm64: GOARCH = arm64
%amd64: GOARCH = amd64

# 后缀变量设置
windows% : EXT = .exe

all:clean build publish
clean:
	@rm -rf bin/${BUILD_NAME}*
# https://golang.org/doc/install/source#environment
build: $(BUILD_TARGET)
# 各目标版本编译
windows_386:
	@$(BUILD_COMMAND)
windows_amd64:
	@$(BUILD_COMMAND)
linux_386:
	@$(BUILD_COMMAND)
linux_amd64:
	@$(BUILD_COMMAND)
linux_arm:
	@$(BUILD_COMMAND)
linux_arm64:
	@$(BUILD_COMMAND)
darwin_386:
	@$(BUILD_COMMAND)
darwin_amd64:
	@$(BUILD_COMMAND)./
darwin_arm:
	@$(BUILD_COMMAND)
darwin_arm64:
	@$(BUILD_COMMAND)
publish:
	@./scripts/publish.sh $(BUILD_NAME) $(GIT_TAG)
ca:
	openssl req -new -nodes -x509 -out conf/server.crt -keyout conf/server.key -days 3650 \
	-subj "/C=DE/ST=NRW/L=Earth/O=Random Company/OU=IT/CN=127.0.0.1/emailAddress=xxxxx@qq.com"
help:
	@echo "make - build and publich"
	@echo "make clean - remove binary file"
	@echo "make gotool - run go tool 'fmt' and 'vet'"
	@echo "make ca - generate ca files"

.PHONY: clean gotool ca help