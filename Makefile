compose-in-memory:
	sudo docker compose --profile in_memory up -d

compose-sql:
	sudo docker compose --profile sql up -d

grpc:
	protoc \
	--go_out=internal/pkg/tinyURL/delivery/server --go_opt=paths=source_relative \
	--go-grpc_out=internal/pkg/tinyUrl/delivery/server --go-grpc_opt=paths=source_relative \
	api/server.proto --proto_path=api

migrate:
	migrate create -ext sql -dir migrations/ -seq init_schema

mock:
	mockgen -source=internal/pkg/tinyURL/delivery/tinyURL.go -destination=internal/pkg/tinyURL/delivery/mocks/tinyURL_mock.go && \
    mockgen -source=internal/pkg/tinyURL/usecase/tinyURL.go -destination=internal/pkg/tinyURL/usecase/mocks/tinyURL_mock.go

test:
	go test ./... -v -coverpkg=./... -coverprofile=cover.out.tmp && cat cover.out.tmp | grep -v "mock.go" | grep -v "pb.go" > cover.out && go tool cover -html=cover.out
