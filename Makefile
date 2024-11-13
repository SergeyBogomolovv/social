.PHONY: gen-proto
gen-proto:
	@protoc --proto_path=pkg/proto \
    --go_out=pkg/proto/generated --go_opt=paths=source_relative \
    --go-grpc_out=pkg/proto/generated --go-grpc_opt=paths=source_relative \
    pkg/proto/*.proto
