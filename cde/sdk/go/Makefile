# Make does not offer a recursive wildcard function, so here's one:
rwildcard=$(wildcard $1$2) $(foreach d,$(wildcard $1*),$(call rwildcard,$d/,$2))

SHELL := /bin/bash
NAME := cde
BUILD_TARGET = build
MAIN_SRC_FILE=main.go
GO := GO111MODULE=on go
GO_NOMOD :=GO111MODULE=off go
REV := $(shell git rev-parse --short HEAD 2> /dev/null || echo 'unknown')
ORG := cdfoundation
ORG_REPO := $(ORG)/$(NAME)
RELEASE_ORG_REPO := $(ORG_REPO)
ROOT_PACKAGE := github.com/$(ORG_REPO)
GO_VERSION := 1.16
GO_DEPENDENCIES := $(call rwildcard,pkg/,*.go) $(call rwildcard,cmd/,*.go)

BRANCH     := $(shell git rev-parse --abbrev-ref HEAD 2> /dev/null  || echo 'unknown')
BUILD_DATE := $(shell date +%Y%m%d-%H:%M:%S)
CGO_ENABLED = 0

REPORTS_DIR=$(BUILD_TARGET)/reports

GOTEST := $(GO) test

# set dev version unless VERSION is explicitly set via environment
VERSION ?= $(shell echo "$$(git for-each-ref refs/tags/ --count=1 --sort=-version:refname --format='%(refname:short)' 2>/dev/null)-dev+$(REV)" | sed 's/^v//')





ifdef DISABLE_TEST_CACHING
GOTEST += -count=1
endif

TEST_PACKAGE ?= ./...
COVER_OUT:=$(REPORTS_DIR)/cover.out
COVERFLAGS=-coverprofile=$(COVER_OUT) --covermode=count --coverpkg=./...

build: $(GO_DEPENDENCIES) clean
	go mod download
	CGO_ENABLED=$(CGO_ENABLED) $(GO) $(BUILD_TARGET) -o build/$(NAME) $(MAIN_SRC_FILE)

test: ## Run tests with the "unit" build tag
	CGO_ENABLED=$(CGO_ENABLED) $(GOTEST) --tags=unit -failfast -short ./...


install: $(GO_DEPENDENCIES) ## Install the binary
	GOBIN=${GOPATH}/bin $(GO) install  $(MAIN_SRC_FILE)

linux: ## Build for Linux
	CGO_ENABLED=$(CGO_ENABLED) GOOS=linux GOARCH=amd64 $(GO) $(BUILD_TARGET) -o build/linux/$(NAME) $(MAIN_SRC_FILE)
	chmod +x build/linux/$(NAME)

arm: ## Build for ARM
	CGO_ENABLED=$(CGO_ENABLED) GOOS=linux GOARCH=arm $(GO) $(BUILD_TARGET) -o build/arm/$(NAME) $(MAIN_SRC_FILE)
	chmod +x build/arm/$(NAME)

win: ## Build for Windows
	CGO_ENABLED=$(CGO_ENABLED) GOOS=windows GOARCH=amd64 $(GO) $(BUILD_TARGET) -o build/win/$(NAME)-windows-amd64.exe $(MAIN_SRC_FILE)

darwin: ## Build for OSX
	CGO_ENABLED=$(CGO_ENABLED) GOOS=darwin GOARCH=amd64 $(GO) $(BUILD_TARGET) -o build/darwin/$(NAME) $(MAIN_SRC_FILE)
	chmod +x build/darwin/$(NAME)

.PHONY: goreleaser
goreleaser:
	step-go-releaser --organisation=$(ORG) --revision=$(REV) --branch=$(BRANCH) --build-date=$(BUILD_DATE) --go-version=$(GO_VERSION) --root-package=$(ROOT_PACKAGE) --version=$(VERSION)

.PHONY: release
release: clean linux test

release-all: release linux win darwin

.PHONY: clean
clean: ## Clean the generated artifacts
	rm -rf build release dist

.PHONY: fmt
fmt: importfmt ## Format the code
	$(eval FORMATTED = $(shell $(GO) fmt ./...))
	@if [ "$(FORMATTED)" == "" ]; \
      	then \
      	    echo "All Go files properly formatted"; \
      	else \
      		echo "Fixed formatting for: $(FORMATTED)"; \
      	fi

.PHONY: importfmt
importfmt: get-fmt-deps
	@echo "Formatting the imports..."
	goimports -w $(GO_DEPENDENCIES)

get-fmt-deps: ## Install test dependencies
	$(GO_NOMOD) get golang.org/x/tools/cmd/goimports

.PHONY: lint
lint: ## Lint the code
	./hack/gofmt.sh
	./hack/linter.sh
	./hack/generate.sh

.PHONY: all
all: fmt build test lint
