NAME := plaid
PREFIX=github.com/tyamell/plaid-go
UNAME=$(shell uname)
GIT_REVISION := $(shell git rev-parse --short HEAD)

ifeq "$(UNAME)" "Darwin"
    BUILD_FLAGS=-ldflags="-s -X plaid.CurrentVersion.Revision=$(GIT_REVISION)"
else
    BUILD_FLAGS=-ldflags="-X plaid.CurrentVersion.Revision=$(GIT_REVISION)"
endif

# Workaround for GO15VENDOREXPERIMENT bug (https://github.com/golang/go/issues/11659)
ALL_PACKAGES=$(shell go list ./... | grep -v /vendor/ | grep -v /scripts)

LINTPKGS := $(shell find . -maxdepth 1 -type d -not -path '*/\.*'|grep -v -e bin -e vendor -e scripts)


.PHONY: coverage vendor build unittest lint test help
.DEFAULT_GOAL := help

run: ## Run the RPC server
	go run main.go


vendor: glide.yaml ## Use glide to vendor dependencies
	@glide up -v
	@glide-vc --use-lock-file --no-tests --only-code

build:  ## Build dev versions
	go build $(BUILD_FLAGS) -o bin/plaid $(PREFIX)/plaid

build-race:  ## Build dev versions with the race flag
	go build -race $(BUILD_FLAGS) -o bin/plaid $(PREFIX)/plaid

# TODO: Release management
# 	https://github.com/goreleaser/goreleaser
coverage.out:
	@ count=1 ;for package in $(ALL_PACKAGES) ; do \
      go test -coverprofile=$$count.cover.out $(BUILD_FLAGS) $$package  ; \
      ((count = count + 1)) ; \
    done
	@echo "mode: set" > coverage.out && cat *.cover.out | grep -v mode: | sort -r >> coverage.out
	@rm  *.cover.out

integrationtests: ## Run all integration tests
	ginkgo -r -p $PREFIX/integration_tests
	
unittest: ## Run all tests
	@ count=1 ;for package in $(ALL_PACKAGES) ; do \
      go test -coverprofile=$$count.cover.out $(BUILD_FLAGS) $$package  ; \
      ((count = count + 1)) ; \
    done
	@echo "mode: set" > coverage.out && cat *.cover.out | grep -v mode: | sort -r >> coverage.out
	@rm  *.cover.out

lint: ## Lint all files
	@sh scripts/run_lints.sh $(LINTPKGS)

test: unittest lint ## Runs unittest and lint


coverage: coverage.out
	go tool cover -o coverage.html -html=coverage.out
	open coverage.html

help:
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'