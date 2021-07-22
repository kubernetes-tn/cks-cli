VERSION := $(shell git describe --tags)
LDFLAGS=-ldflags "-s -w -X=main.version=$(VERSION)"

.PHONY: build
build:
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) .