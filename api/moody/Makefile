.PHONY: generate regenerate

PROTO_FILES := $(wildcard *.proto)
GENERATED_GO := $(patsubst %.proto,%.pb.go,$(PROTO_FILES))
GENERATED_GRPC := $(patsubst %.proto,%_grpc.pb.go,$(PROTO_FILES))

generate: $(GENERATED_GO) $(GENERATED_GRPC)

$(GENERATED_GO) $(GENERATED_GRPC): $(PROTO_FILES)
	@go generate .

regenerate:
	touch $(PROTO_FILES)
	@$(MAKE) generate

clean:
	@rm -f *.pb.go
