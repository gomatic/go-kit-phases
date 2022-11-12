NAME ?= moody

API := $(NAME)

export GOOS ?= $(shell go env GOOS)
export GOARCH ?= $(shell go env GOARCH)
export CGO_ENABLED ?= 0

BINARY ?= bin/$(NAME)-$(GOOS)-$(GOARCH)

CONFIG_PACKAGE := versioning
CONFIG := $(GIT_HOST)/$(GIT_ORG)/$(CONFIG_PACKAGE)

GO_FILES := $(patsubst ./%,%,$(shell find . -type f -name '*.go' -o -type d \( -path ./cmd -o -path ./vendor \) -prune -false))
GO_CLI_FILES := $(shell find cmd/$(API) -name '*.go')

GO_LINTER_COMMAND := $(shell command -v golangci-lint || echo echo missing golangci-lint)
GO_LINTER := $(GO_LINTER_COMMAND) run --allow-parallel-runners --skip-dirs-use-default --disable=typecheck

GO_FORMATTER := $(shell command -v goimports || echo echo install goimports --)

GO_FLAGS := -mod=vendor
GO_LDFLAGS := "-X '$(CONFIG).commit=$(GIT_COMMIT)' -X '$(CONFIG).tag=$(CANDIDATE_TAG)' -X '$(CONFIG).date=$(GIT_DATETIME)' -X '$(CONFIG).build=$(BUILD_NUMBER)' -X '$(CONFIG).branch=$(GIT_BRANCH)'"
GO_BUILD := go build $(GO_FLAGS) -ldflags=$(GO_LDFLAGS)

ifdef DEBUGGING
GO_DEBUG_FLAGS := -gcflags='all=-N -l'
BINARY := $(BINARY)-debug
endif

$(BINARY): $(GO_FILES)
	$(GO_BUILD) $(GO_DEBUG_FLAGS) -o $@ ./cmd/$(NAME)

lint: vet ## Vet and lint
	@$(GO_LINTER)

vet: ## Vet
	go vet $(GO_FLAGS) $(GOVET_FLAGS) ./...

fmt: ## Format
	@$(GO_FORMATTER) -w $(GO_FILES) $(GO_CLI_FILES)
