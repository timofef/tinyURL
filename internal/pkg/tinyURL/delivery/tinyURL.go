package delivery

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	server "github.com/timofef/tinyURL/internal/pkg/tinyURL/delivery/server"
	"github.com/timofef/tinyURL/internal/tinyURL/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type IUsecase interface {
	Add(fullUrl string) (string, error)
	Get(tinyUrl string) (string, error)
}

type TinyUrlHandler struct {
	Usecase IUsecase
	server.UnimplementedTinyUrlServiceServer
}

func (h *TinyUrlHandler) Add(ctx context.Context, fullUrl *server.FullUrl) (*server.TinyUrl, error) {
	err := validation.Validate(fullUrl.Val,
		validation.Required,
		is.URL,
	)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid URL")
	}

	tinyUrl, err := h.Usecase.Add(fullUrl.Val)
	if err != nil {
		return nil, status.Error(codes.Internal, "Server error: "+err.Error())
	}

	logger.MainLogger.Info("Added " + fullUrl.Val + " as " + tinyUrl)

	return &server.TinyUrl{Val: tinyUrl}, nil
}

func (h *TinyUrlHandler) Get(ctx context.Context, tinyUrl *server.TinyUrl) (*server.FullUrl, error) {
	err := validation.Validate(tinyUrl.Val,
		validation.Required,
		is.URL,
	)
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

	logger.MainLogger.Info("Found " + tinyUrl.Val + " as " + fullUrl)

	return &server.FullUrl{Val: fullUrl}, nil
}
