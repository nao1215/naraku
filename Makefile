.PHONY: build test clean

APP         = naraku
VERSION     = $(shell git describe --tags --abbrev=0)
GIT_REVISION := $(shell git rev-parse HEAD)
GO          = go
GO_BUILD    = $(GO) build
GO_TEST     = $(GO) test -v
GO_TOOL     = $(GO) tool
GOOS        = ""
GOARCH      = ""
GO_PKGROOT  = ./...
GO_PACKAGES = $(shell $(GO_LIST) $(GO_PKGROOT))
GO_LDFLAGS  = -ldflags '-X github.com/nao1215/naraku/version.Version=${VERSION}'

build:  ## Build binary
	env GO111MODULE=on GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO_BUILD) $(GO_LDFLAGS) -o $(APP) main.go

server: ## Run server
	env GO111MODULE=on GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) run main.go

generate: ## Generate file automatically
	#docker-compose up -d db
	$(GO) generate ./...
	swag fmt && swag init --output ../docs
	#sqlc generate --file server/app/schema/sqlc.yml 
	#tbls doc --force 
	wire ./...

clean: ## Clean project
	-rm -rf $(APP) cover.out cover.html

test: ## Start unit test for server
	env GOOS=$(GOOS) $(GO_TEST) -cover $(GO_PKGROOT) -coverprofile=coverage.out
	$(GO_TOOL) cover -html=coverage.out -o coverage.html
	
.DEFAULT_GOAL := help
help: ## Show help message  
	@grep -E '^[0-9a-zA-Z_-]+[[:blank:]]*:.*?## .*$$' $(MAKEFILE_LIST) | sort \
	| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[1;32m%-15s\033[0m %s\n", $$1, $$2}'