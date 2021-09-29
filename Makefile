LANG := en_US.UTF-8
SHELL := /bin/bash
.SHELLFLAGS := --norc --noprofile -e -u -o pipefail -c
.DEFAULT_GOAL := build

name := kkn.fi/base62

GOIMPORTS := $(GOPATH)/bin/goimports
STATICCHECK := $(GOPATH)/bin/staticcheck
GOLANGCI-LINT := $(GOPATH)/bin/golangci-lint

.PHONY: build
build:
	go build $(name)

.PHONY: test
test:
	go test $(name)

$(GOIMPORTS):
	go install golang.org/x/tools/cmd/goimports@latest

$(GOLANGCI-LINT):
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.41.1

$(STATICCHECK):
	go install honnef.co/go/tools/cmd/staticcheck@latest

.PHONY: fmt
fmt:
	gofmt -w -s .

.PHONY: goimports
goimports: fmt $(GOIMPORTS)
	$(GOIMPORTS) -w .

.PHONY: staticcheck
staticcheck: $(STATICCHECK)
	$(STATICCHECK) -go 1.17 ./...

.PHONY: golangci-lint
golangci-lint: $(GOLANGCI-LINT)
	$(GOLANGCI-LINT) run ./...

.PHONY: cover
cover:
	go test -coverprofile=coverage.out $(name)
	go tool cover -html=coverage.out
	@rm -f coverage.out

.PHONY: heat
heat:
	go test -covermode=count -coverprofile=count.out $(name)
	go tool cover -html=count.out
	@rm -f count.out
