NAME := go-metaforce
SRCS := $(shell find . -type d -name vendor -prune -o -type f -name "*.go" -print)
VERSION := 0.1.0
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := -ldflags="-s -w -X \"main.Version=$(VERSION)\" -X \"main.Revision=$(REVISION)\"" 
DIST_DIRS := find * -type d -exec

.DEFAULT_GOAL := bin/$(NAME) 

.PHONY: test
test:
	@go test -cover

.PHONY: clean
clean:
	@rm -rf bin/*
	@rm -rf vendor/*
	@rm -rf dist/*

.PHONY: format
format: import
	-@goimports -w $(SRCS)
	@gofmt -w $(SRCS)

.PHONY: import
import:
	go get golang.org/x/tools/cmd/goimports

.PHONY: dep
dep:
ifeq ($(shell command -v dep 2> /dev/null),)
	curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
endif

.PHONY: deps
deps:
	dep ensure

.PHONY: dist
docker-build:
	docker build . -t $(NAME)

.PHONY: build
build:
	go build .

.PHONY: run
run:
	go run main/main.go
