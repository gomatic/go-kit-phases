.PHONY: generate regenerate

include make/env.mk
include make/go.mk

build: $(BINARY)

linux darwin windows:
	$(MAKE) build GOOS=$@ GOARCH=amd64

deps:
	go mod vendor
	go mod tidy -compat=1.17

generate regenerate:
	cd api/$(API); $(MAKE) $@

clean:
	cd api/$(API); $(MAKE) $@
