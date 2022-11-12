GRPC_PORT ?= 8091
HTTP_PORT ?= 8093
JRPC_PORT ?= 8095

export JWT_SECRET := xyz

THIS := $(abspath $(firstword $(MAKEFILE_LIST)))
THISD := $(patsubst %/,%,$(dir $(THIS)))

start:
	$(THISD)/../bin/moody-$(GOOS)-$(GOARCH) server

grpc-cli-%:: PORT ?= $(GRPC_PORT)
grpc-cli-%:
	grpc_cli $(patsubst grpc-cli-%,%,$@) localhost:$(PORT)

grpcurl: PORT ?= $(GRPC_PORT)
grpcurl: RESOURCE ?= self
grpcurl: METHOD ?= Get
grpcurl:
	grpcurl -plaintext localhost:$(PORT) storage.$(RESOURCE).$(METHOD)

test-grpc: PORT ?= $(GRPC_PORT)
test-grpc: RESOURCE ?= self
test-grpc:
	curl \
			-X POST \
		localhost:$(PORT)/moody

#

test-http-self:
	$(MAKE) -f $(THIS) test-http RESOURCE=$(patsubst test-http-%,%,$@)

var/token-$(JWT_SECRET): $(THISD)/../var/token-$(JWT_SECRET)

$(THISD)/../var/token-$(JWT_SECRET): $(THISD)/../var
	$(THISD)/../bin/moody-$(GOOS)-$(GOARCH) client test token | tee $@

$(THISD)/../var:
	mkdir -p $@

test-http: PORT ?= $(HTTP_PORT)
test-http: RESOURCE ?= self
test-http: var/token-$(JWT_SECRET)
	curl \
			-X GET \
			-H 'Authorization: Bearer $(shell cat $(THISD)/../var/token-$(JWT_SECRET))' \
			-H 'Content-Type: application/json' \
		localhost:$(PORT)/moody/$(RESOURCE)

#

test-jsonrpc-self:
	$(MAKE) -f $(THIS) test-jsonrpc RESOURCE=$(patsubst test-jsonrpc-%,%,$@)

test-jsonrpc: PORT ?= $(JRPC_PORT)
test-jsonrpc: RESOURCE ?= self
test-jsonrpc: METHOD ?= retrieve
test-jsonrpc: var/token-$(JWT_SECRET)
	curl \
			-X POST \
			-H 'Authorization: Bearer $(shell cat $(THISD)/../var/token-$(JWT_SECRET))' \
			-H 'Content-Type: application/json' \
			-d '{"jsonrpc":"2.0","id":"id","method":"$(RESOURCE)/$(METHOD)","params":{"path":"/"}}' \
		localhost:$(PORT)

