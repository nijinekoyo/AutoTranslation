GO := go

APPNAME := AutoTranslation
VERSION := $(shell git describe --tags --abbrev=0)
BUILD_DIR := build

.PHONY: all
all: build

# 安装依赖
.PHONY: deps
deps:
	$(GO) mod tidy

# 编译
.PHONY: build
OUTPUT_PATH := ${BUILD_DIR}/$(APPNAME)-$(VERSION)
ifeq ($(OS),Windows_NT)
	OUTPUT_PATH := ${OUTPUT_PATH}.exe
endif
build:
	$(GO) build \
	-o $(OUTPUT_PATH) \
	.

# 运行
.PHONY: run
run:
	$(GO) run .
