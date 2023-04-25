compose-in-memory:
	sudo docker compose --profile in_memory up -d

compose-sql:
	sudo docker compose --profile sql up -d

grpc:
	protoc \
	--go_out=internal/delivery/server --go_opt=paths=source_relative \
	--go-grpc_out=internal/delivery/server --go-grpc_opt=paths=source_relative \
	api/server.proto --proto_path=api

migrate:
	migrate create -ext sql -dir migrations/ -seq init_schema

mock:
	mockgen -source=internal/delivery/tinyURL.go -destination=internal/delivery/mocks/tinyURL_mock.go && \
    mockgen -source=internal/usecase/tinyURL.go -destination=internal/usecase/mocks/tinyURL_mock.go

test:
	go test ./... -v -coverpkg=./... -coverprofile=cover.out.tmp && \
	cat cover.out.tmp | grep -v "mock.go" | grep -v "pb.go" > cover.out && go tool cover -html=cover.out

