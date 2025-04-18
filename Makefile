SHELL = /usr/bin/env bash
.SHELLFLAGS = -ecuo pipefail
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules
MAKEFLAGS += --no-builtin-variables
.DEFAULT_GOAL := help

# ==============================================================================
# Variables
# ==============================================================================
CI ?=
OUTPUTDIR ?= output
COVERPROFILE ?= coverage.out
COVERHTML ?= coverage.html

# ==============================================================================
# Functions
# ==============================================================================
define open_browser
	@case $$(uname -s) in \
		Linux) xdg-open $(1) ;; \
		Darwin) open $(1) ;; \
		*) echo "Unsupported platform" ;; \
	esac
endef

# ==============================================================================
# Targets
# ==============================================================================

## Dependencies:
.PHONY: install
install: ## Install dependencies
	go mod download -v
	go mod verify

.PHONY: update
update: ## Update dependencies
	go get -u ./...
	go mod tidy -v

## Code Quality:
.PHONY: fmt
fmt: ## Format code
	golangci-lint fmt --verbose

.PHONY: lint
lint: ## Run linter
	golangci-lint run --verbose

## Testing:
.PHONY: test
test: COVER_PROFILE_PATH = $(OUTPUTDIR)/$(COVERPROFILE)
test: COVER_HTML_PATH = $(OUTPUTDIR)/$(COVERHTML)
test: ## Run tests
	mkdir -pv $(dir $(COVER_PROFILE_PATH))
	go test -v \
		-outputdir=$(dir $(COVER_PROFILE_PATH)) \
		-coverprofile=$(notdir $(COVER_PROFILE_PATH)) \
		-coverpkg=./... \
		-run= ./... | \
		tee $(OUTPUTDIR)/test.log
ifeq ($(CI),)
	mkdir -pv $(dir $(COVER_HTML_PATH))
	go tool cover -html=$(COVER_PROFILE_PATH) -o $(COVER_HTML_PATH)
	@echo "🌐 Run 'make browser/cover' to open coverage report in browser"
endif

.PHONY: browser/cover
browser/cover: ## Open browser with Go coverage report
	$(call open_browser,$(OUTPUTDIR)/$(COVERHTML))

## Build:
.PHONY: build
build: ## Build the project
	goreleaser build --verbose --snapshot --clean --timeout=5m

.PHONY: clean
clean: ## Clean up generated files
	rm -rfv $(OUTPUTDIR) || true

## Help:
.PHONY: help
help: GREEN  := $(shell tput -Txterm setaf 2)
help: YELLOW := $(shell tput -Txterm setaf 3)
help: CYAN   := $(shell tput -Txterm setaf 6)
help: RESET  := $(shell tput -Txterm sgr0)
help: ## Show this help
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} { \
		if (/^[a-zA-Z0-9_\/-]+:.*?##.*$$/) {printf "    ${YELLOW}%-20s${GREEN}%s${RESET}\n", $$1, $$2} \
		else if (/^## .*$$/) {printf "  ${CYAN}%s${RESET}\n", substr($$1,4)} \
		}' $(MAKEFILE_LIST)
