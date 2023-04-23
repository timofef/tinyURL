compose-in-memory:
	sudo docker compose --profile in_memory up -d

compose-postgres:
	sudo docker compose --profile sql up -d

grpc:
	protoc \
	--go_out=internal/pkg/tinyURL/delivery/server --go_opt=paths=source_relative \
	--go-grpc_out=internal/pkg/tinyUrl/delivery/server --go-grpc_opt=paths=source_relative \
	api/server.proto --proto_path=api

migrate:
	migrate create -ext sql -dir migrations/ -seq init_schema

mock:
	mockgen \
	-source=internal/pkg/tinyURL/usecase/tinyURL.go \
	-destination=internal/pkg/tinyURL/repository/mocks/tinyURL_mock.go