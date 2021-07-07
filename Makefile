GOCMD=$(shell which go)
PKGBASE=github.com/erumble
PROJECTNAME=$(shell basename "$(PWD)")
BINARY_NAME=echo
BUILD_VERSION?=0.0.0
EXPORT_RESULT?=false # for CI please set EXPORT_RESULT to true

GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
CYAN   := $(shell tput -Txterm setaf 6)
RESET  := $(shell tput -Txterm sgr0)

.PHONY: all build vendor

all: help

## Build:
build: ## Build your project and put the output binary in bin/
	mkdir -p bin
	CGO_ENABLED=0 $(GOCMD) build -ldflags "-X $(PKGBASE)/$(PROJECTNAME)/pkg/cli.version=$(BUILD_VERSION)" -o bin/$(BINARY_NAME) $(PKGBASE)/$(PROJECTNAME)/cmd/$(BINARY_NAME)

clean: ## Remove build related file
	rm -rf ./bin
	rm -f ./checkstyle-report.xml

vendor: ## Copy of all packages needed to support builds and tests in the vendor directory
	$(GOCMD) mod vendor
	$(GOCMD) mod tidy

## Lint:
lint: lint-go ## Run all available linters

lint-go: ## Use golintci-lint on your project
	$(eval OUTPUT_OPTIONS = $(shell [ "${EXPORT_RESULT}" == "true" ] && echo "--out-format checkstyle ./... | tee /dev/tty > checkstyle-report.xml" || echo "" ))
	docker run --rm -v $(shell pwd):/app -w /app golangci/golangci-lint:latest-alpine golangci-lint run --deadline=65s $(OUTPUT_OPTIONS)

## Help:
help: ## Show this help.
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} { \
		if (/^[a-zA-Z_-]+:.*?##.*$$/) {printf "    ${YELLOW}%-20s${GREEN}%s${RESET}\n", $$1, $$2} \
		else if (/^## .*$$/) {printf "  ${CYAN}%s${RESET}\n", substr($$1,4)} \
		}' $(MAKEFILE_LIST)