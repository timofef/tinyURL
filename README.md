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
### Makefile
PostgreSQL:
```
todo
```
In-memory:
```
todo
```
### Вручную
PostgreSQL
```
todo
```
In-memory:
```
todo
```
## Использованные команды для генерации файлов
Mock:
```
```
gRPC:
```
protoc --go_out=internal/pkg/tinyURL/delivery/server --go_opt=paths=source_relative --go-grpc_out=internal/pkg/tinyUrl/delivery/server --go-grpc_opt=paths=source_relative internal/pkg/tinyURL/delivery/server/proto/server.proto --proto_path=internal/pkg/tinyURL/delivery/server/proto
```
Миграции:
```
migrate create -ext sql -dir migrations/ -seq init_schema
```