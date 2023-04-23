package delivery

import (
	"context"
	server "github.com/timofef/tinyURL/internal/pkg/tinyURL/delivery/server"
	"github.com/timofef/tinyURL/internal/pkg/tinyURL/usecase"
	"github.com/timofef/tinyURL/internal/tinyURL/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/url"
)

type TinyUrlHandler struct {
	Server  server.TinyUrlServiceServer
	Usecase usecase.IUsecase
	server.UnimplementedTinyUrlServiceServer
}

func (h *TinyUrlHandler) Add(ctx context.Context, fullUrl *server.FullUrl) (*server.TinyUrl, error) {
	// TODO fix
	_, err := url.ParseRequestURI(fullUrl.Val)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid URL")
	}

	tinyUrl, err := h.Usecase.Add(fullUrl.Val)
	if err != nil {
		return nil, status.Error(codes.Internal, "Server error: "+err.Error())
	}

	logger.MainLogger.LogInfo("Added " + fullUrl.Val + " as " + tinyUrl)

	return &server.TinyUrl{Val: tinyUrl}, nil
}

func (h *TinyUrlHandler) Get(ctx context.Context, tinyUrl *server.TinyUrl) (*server.FullUrl, error) {
	// TODO fix
	_, err := url.ParseRequestURI(tinyUrl.Val)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid URL")
	}

	fullUrl, err := h.Usecase.Get(tinyUrl.Val)
	if err != nil {
		return nil, status.Error(codes.Internal, "Server error: "+err.Error())
	}
	if fullUrl == "" {
		return nil, status.Error(codes.NotFound, "Can't find URL: "+tinyUrl.Val)
	}

	logger.MainLogger.LogInfo("Found " + tinyUrl.Val + " as " + fullUrl)

	return &server.FullUrl{Val: fullUrl}, nil
}
