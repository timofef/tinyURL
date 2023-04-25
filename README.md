# tinyURL

Реализовать сервис, предоставляющий API по созданию сокращённых ссылок.

Ссылка должна быть:

* уникальной; на один оригинальный URL должна ссылаться только одна сокращенная ссылка;

* длиной 10 символов;

* из символов латинского алфавита в нижнем и верхнем регистре, цифр и символа _ (подчеркивание).


Сервис должен быть написан на Go и принимать следующие запросы по http:
1. Метод Post, который будет сохранять оригинальный URL в базе и возвращать сокращённый.
2. Метод Get, который будет принимать сокращённый URL и возвращать оригинальный.

**Условие со звёздочкой:** cделать работу сервиса через GRPC, то есть составить proto и реализовать сервис с
двумя соответствующими эндпойнтами.

Решение должно соответствовать условиям:

* сервис распространён в виде Docker-образа;

* в качестве хранилища ожидаем in-memory решение и PostgreSQL. Какое хранилище использовать, указывается параметром при запуске сервиса;

* реализованный функционал покрыт Unit-тестами.

## Запуск приложения

### PostgreSQL:

```
make compose-sql
```
или
```
sudo docker compose --profile sql up -d
```

### In-memory:

```
make compose-in-memory
```
или
```
sudo docker compose --profile in_memory up -d
```

## Использованные команды для генерации файлов
### Mock:
```
make mock
```
или
```
mockgen -source=internal/delivery/tinyURL.go -destination=internal/delivery/mocks/tinyURL_mock.go && mockgen -source=internal/usecase/tinyURL.go -destination=internal/usecase/mocks/tinyURL_mock.go
```
### gRPC:
```
make grpc
```
или
```
protoc --go_out=internal/delivery/server --go_opt=paths=source_relative --go-grpc_out=internal/delivery/server --go-grpc_opt=paths=source_relative api/server.proto --proto_path=api
```
### Миграции:
```
make migrate
```
или
```
migrate create -ext sql -dir migrations/ -seq init_schema
```
## Тесты

Запуск тестов и генерация html-файла с покрытием:
```
make test
```
или
```
go test ./... -v -coverpkg=./... -coverprofile=cover.out.tmp && cat cover.out.tmp | grep -v "mock.go" | grep -v "pb.go" > cover.out && go tool cover -html=cover.out
```