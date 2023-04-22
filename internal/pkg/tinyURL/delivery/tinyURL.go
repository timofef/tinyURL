package delivery

import (
	"context"
	server "github.com/timofef/tinyURL/internal/pkg/tinyURL/delivery/server"
	"github.com/timofef/tinyURL/internal/pkg/tinyURL/usecase"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TinyUrlHandler struct {
	Server  server.TinyUrlServiceServer
	Usecase usecase.IUsecase
}

func (h *TinyUrlHandler) Add(ctx context.Context, fullUrl *server.FullUrl) (*server.TinyUrl, error) {
	// TODO: check if is url
	ok := true

	if !ok {
		return nil, status.Error(codes.InvalidArgument, "Invalid URL")
	}

	// create
	tinyUrl, err := h.Usecase.Add(fullUrl.Val)
	if err != nil {
		return nil, status.Error(codes.Internal, "Server error: "+err.Error())
	}

	return &server.TinyUrl{Val: tinyUrl}, nil
}

func (h *TinyUrlHandler) Get(ctx context.Context, tinyUrl *server.TinyUrl) (*server.FullUrl, error) {
	// TODO: check if is url
	ok := true

	if !ok {
		return nil, status.Error(codes.InvalidArgument, "Invalid URL")
	}

	fullUrl, err := h.Usecase.Get(tinyUrl.Val)
	if err != nil {
		return nil, status.Error(codes.Internal, "Server error: "+err.Error())
	}
	if fullUrl == "" {
		return nil, status.Error(codes.NotFound, "Can't find URL: " + tinyUrl.Val)
	}

	return &server.FullUrl{Val:}, nil
}
