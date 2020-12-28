.PHONY: generate regenerate

GO_FILES := $(shell find api internal -name '*.go')
SERVER_GO_FILES := $(shell find cmd/server -name '*.go')
CLIENT_GO_FILES := $(shell find cmd/client -name '*.go')

build: bin/server bin/client

bin/server: $(GO_FILES) $(SERVER_GO_FILES)
	go build -o $@ ./cmd/server

bin/client: $(GO_FILES) $(CLIENT_GO_FILES)
	go build -o $@ ./cmd/server

deps:
	go mod vendor
	go mod tidy

generate regenerate:
	cd api/moody; $(MAKE) $@

clean:
	cd api/moody; $(MAKE) $@
