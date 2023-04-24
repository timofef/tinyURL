# tinyURL

Реализовать сервис, предоставляющий API по созданию сокращённых ссылок.

Ссылка должна быть:

>— уникальной; на один оригинальный URL должна ссылаться только одна сокращенная ссылка;

>— длиной 10 символов;

>— из символов латинского алфавита в нижнем и верхнем регистре, цифр и символа _ (подчеркивание).

Сделать работу сервиса через GRPC, то есть составить proto и реализовать сервис с
двумя соответствующими эндпойнтами.

Решение должно соответствовать условиям:

>— сервис распространён в виде Docker-образа;

>— в качестве хранилища ожидаем in-memory решение и PostgreSQL. Какое хранилище использовать, указывается параметром при запуске сервиса;

>— реализованный функционал покрыт Unit-тестами.

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
mockgen -source=internal/pkg/tinyURL/delivery/tinyURL.go -destination=internal/pkg/tinyURL/delivery/mocks/tinyURL_mock.go && mockgen -source=internal/pkg/tinyURL/usecase/tinyURL.go -destination=internal/pkg/tinyURL/usecase/mocks/tinyURL_mock.go
```
### gRPC:
```
make grpc
```
или
```
protoc --go_out=internal/pkg/tinyURL/delivery/server --go_opt=paths=source_relative --go-grpc_out=internal/pkg/tinyUrl/delivery/server --go-grpc_opt=paths=source_relative api/server.proto --proto_path=api
```
### Миграции:
```
make migrate
```
или
```
migrate create -ext sql -dir migrations/ -seq init_schema
```