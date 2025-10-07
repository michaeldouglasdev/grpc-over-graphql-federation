.PHONY: start generate

all: generate

start:
	$(MAKE) service \
	wait

service:
	go run ./cmd/service/main.go

client:
	go run ./cmd/client/main.go

generate-proto:
	npx wgc@latest grpc-service generate -i ./pkg/graph/schema.graphql -o ./pkg/proto/service/v1 --go-package "github.com/wundergraph/service/pkg/generated/service/v1" Service

generate-buf:
	buf generate

generate: generate-proto generate-buf