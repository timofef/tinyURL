package delivery

import (
	"context"
	server "github.com/timofef/tinyURL/internal/pkg/tinyURL/delivery/server"
)

type IHandler interface {
	Add(ctx context.Context, fullUrl *server.FullUrl) (*server.TinyUrl, error)
	Get(ctx context.Context, tinyUrl *server.TinyUrl) (*server.FullUrl, error)
}
