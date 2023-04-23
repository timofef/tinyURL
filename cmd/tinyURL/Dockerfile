FROM golang:1.20

ADD . /app
WORKDIR /app
RUN go build ./cmd/tinyURL/main.go

EXPOSE 5555

ENTRYPOINT [ "./main" ]