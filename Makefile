SHELL = /bin/bash
export PATH := $(abspath bin/):${PATH}

ifeq ($(strip $(shell git status --porcelain 2>/dev/null)),)
  export GIT_TREE_STATE=clean
else
  export GIT_TREE_STATE=dirty
endif

export GO111MODULE := on
export GOPROXY = https://proxy.golang.org,direct

all: clean setup fmt lint test

# Install all the build and lint dependencies
setup:
	curl -sfL https://install.goreleaser.com/github.com/goreleaser/goreleaser.sh | sh
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh
.PHONY: setup

.PHONY: test
test:
	@echo "==> Testing all packages"
	@go test -v -failfast -race -coverpkg=./... -covermode=atomic -coverprofile=profile.out ./... -run . -timeout=2m

.PHONY: cover
cover: test
	@echo "==> Coverage all packages"
	@go tool cover -html=profile.out


.PHONY: lint
lint: bin/golangci-lint ## Run linter
	@echo "==> Run golangci-lint"
	bin/golangci-lint run

.PHONY: fix
fix: bin/golangci-lint ## Fix lint violations
	@echo "==> Run golangci-lint fix"
	bin/golangci-lint run --fix

.PHONY: fmt
fmt:
	@echo "==> fmt all go files"
	@go fmt ./...


.PHONY: snapshot
snapshot:
	@echo "==> snapshot"
	bin/goreleaser --snapshot --skip-publish --rm-dist


.PHONY: build
build:
	@echo "==> build"
	bin/goreleaser


.PHONY: clean
clean:
	@echo "==> Cleaning releases"
	@go clean -i -x ./...
