ROOT = $(shell pwd)
SERVICE_NAME = $(shell basename "$(PWD)")
GO ?= go
OS = $(shell uname -s | tr A-Z a-z)
export GOBIN = ${ROOT}/bin

LINT = ${GOBIN}/golangci-lint
LINT_DOWNLOAD = curl --progress-bar -SfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s latest
GOPLANTUML = ${GOBIN}/goplantuml
GOPLANTUML_DOWNLOAD = $(GO) get github.com/jfeliu007/goplantuml/cmd/goplantuml
VERSION_TAG = $(shell git describe --tags --abbrev=0 --always)
VERSION_COMMIT = $(shell git rev-parse --short HEAD)
VERSION_DATE = $(shell git show -s --format=%cI HEAD)
VERSION = -X main.versionTag=$(VERSION_TAG) -X main.versionCommit=$(VERSION_COMMIT)  -X main.versionDate=$(VERSION_DATE) -X main.serviceName=$(SERVICE_NAME)
TPARSE = $(GOBIN)/tparse
TPARSE_DOWNLOAD = $(GO) get github.com/mfridman/tparse
COMPILEDEAMON = $(GOBIN)/CompileDaemon
COMPILEDEAMON_DOWNLOAD = $(GO) get github.com/githubnemo/CompileDaemon
MIGRATE = ${GOBIN}/migrate
MIGRATE_DOWNLOAD = (curl --progress-bar -fL -o $(MIGRATE).tar.gz https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.$(OS)-amd64.tar.gz; tar -xzvf $(MIGRATE).tar.gz -C $(GOBIN); mv $(MIGRATE).$(OS)-amd64 $(MIGRATE); rm $(MIGRATE).tar.gz)
MIGRATE_CONFIG = -source file://migrations -database "mysql://${MYSQL_USER}:${MYSQL_PASS}@tcp(${MYSQL_HOST}:${MYSQL_PORT})/${MYSQL_DBNAME}"
PATH := $(PATH):$(GOBIN)

.PHONY: help
help: ## Display this help message
	@ cat $(MAKEFILE_LIST) | grep -e "^[a-zA-Z_\-]*: *.*## *" | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: build
build: ## Build development binary file
	@ $(GO) build -race -ldflags '$(VERSION)' -o ./bin/${SERVICE_NAME} ./cmd/...

.PHONY: run
run: ## run as development reload if code changes referrer-ip=referrer-ip referrer-port=referrer-port
	@ test -e $(COMPILEDEAMON) || $(COMPILEDEAMON_DOWNLOAD)
	@ $(COMPILEDEAMON) --build="make build" --command="$(GOBIN)/$(SERVICE_NAME)"

.PHONY: mod
mod: ## Get dependency packages
	@ $(GO) mod tidy

.PHONY: test
test: ## Run data race detector
	@ test -e $(TPARSE) || $(TPARSE_DOWNLOAD)
	@ $(GO) test -timeout 1000s -short -race ./... -json -cover | $(TPARSE) -all -smallscreen

.PHONY: coverage
coverage: ## check coverage test code of sample https://penkovski.com/post/gitlab-golang-test-coverage/
	@ $(GO) test -timeout 1000s ./... -coverprofile=coverage.out
	@ $(GO) tool cover -func=coverage.out
	@ $(GO) tool cover -html=coverage.out -o coverage.html;

.PHONY: lint
lint: ## Lint the files
	@ test -e $(LINT) || $(LINT_DOWNLOAD)
	@ $(LINT) version
	@ $(LINT) --timeout 10m run

.PHONY: uml
uml: ## Create UML diagram in diagram.puml file
	@ test -e $(GOPLANTUML) || $(GOPLANTUML_DOWNLOAD)
	@ $(GOPLANTUML) -recursive . > diagram.puml

.PHONY: migrate
migrate: ## base migrate
	@ test -e $(MIGRATE) || $(MIGRATE_DOWNLOAD)
	@ $(MIGRATE) --version

.PHONY: migrate-up
migrate-up:migrate ## Apply all up migrations
	@ $(MIGRATE) $(MIGRATE_CONFIG) up

.PHONY:	migrate-down
migrate-down:migrate ## Apply all down migrations
	@ $(MIGRATE) $(MIGRATE_CONFIG) down

.PHONY: migrate-drop
migrate-drop:migrate ## Apply all down migrations
	@ $(MIGRATE) $(MIGRATE_CONFIG) drop

.PHONY: env
env: ## create env file from .env.example and read env file & export to terminal
	@ test -e ./.env && echo ./.env exists || cp ./.env.example ./.env
	@ export $(grep -v '^#' .env | xargs -d '\n')
