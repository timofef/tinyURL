grpc:
	protoc \
	--go_out=internal/pkg/tinyURL/delivery/server --go_opt=paths=source_relative \
	--go-grpc_out=internal/pkg/tinyUrl/delivery/server --go-grpc_opt=paths=source_relative internal/pkg/tinyURL/delivery/server/proto/server.proto \
	--proto_path=internal/pkg/tinyURL/delivery/server/proto