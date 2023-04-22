# tinyURL

Реализовать сервис, предоставляющий API по созданию сокращённых ссылок.

Ссылка должна быть:

>— Уникальной; на один оригинальный URL должна ссылаться только одна сокращенная ссылка;

>— Длиной 10 символов;

>— Из символов латинского алфавита в нижнем и верхнем регистре, цифр и символа _ (подчеркивание).

Сделать работу сервиса через GRPC, то есть составить proto и реализовать сервис с
двумя соответствующими эндпойнтами.

Решение должно соответствовать условиям:

>— Сервис распространён в виде Docker-образа;

>— В качестве хранилища ожидаем in-memory решение и PostgreSQL. Какое хранилище использовать, указывается параметром при запуске сервиса;

>— Реализованный функционал покрыт Unit-тестами.

## Запуск

### В контейнере

С PostgreSQL:

```
todo
```

In-memory:
```
todo
```

### Локально

С PostgreSQL:

```
todo
```

In-memory:
```
todo
```

## Генерация файлов

### Через Makefile
Mock:
```
```

gRPC:
```
make grpc
```

### Вручную
Mock:
```
```

gRPC:
```
protoc --go_out=internal/pkg/tinyURL/delivery/server --go_opt=paths=source_relative --go-grpc_out=internal/pkg/tinyUrl/delivery/server --go-grpc_opt=paths=source_relative internal/pkg/tinyURL/delivery/server/proto/server.proto --proto_path=internal/pkg/tinyURL/delivery/server/proto
```