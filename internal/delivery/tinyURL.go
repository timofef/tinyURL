package delivery

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/timofef/tinyURL/internal/delivery/server"
	"github.com/timofef/tinyURL/internal/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type IUsecase interface {
	Add(fullUrl string) (string, error)
	Get(tinyUrl string) (string, error)
}

type TinyUrlHandler struct {
	usecase IUsecase
	server_proto.UnimplementedTinyUrlServiceServer
}

func InitTinyUrlHandler(uc IUsecase) *TinyUrlHandler {
	return &TinyUrlHandler{usecase: uc}
}

func (h *TinyUrlHandler) Add(ctx context.Context, fullUrl *server_proto.FullUrl) (*server_proto.TinyUrl, error) {
	err := validation.Validate(
		fullUrl.Val,
		validation.Required,
		is.URL,
	)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid URL")
	}

	tinyUrl, err := h.usecase.Add(fullUrl.Val)
	if err != nil {
		return nil, status.Error(codes.Internal, "Server error: "+err.Error())
	}

	logger.MainLogger.Info("Added " + fullUrl.Val + " as " + tinyUrl)

	return &server_proto.TinyUrl{Val: tinyUrl}, nil
}

func (h *TinyUrlHandler) Get(ctx context.Context, tinyUrl *server_proto.TinyUrl) (*server_proto.FullUrl, error) {
	err := validation.Validate(tinyUrl.Val,
		validation.Required,
		is.URL,
	)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid URL")
	}

	fullUrl, err := h.usecase.Get(tinyUrl.Val)
	if err != nil {
		return nil, status.Error(codes.Internal, "Server error: "+err.Error())
	}
	if fullUrl == "" {
		return nil, status.Error(codes.NotFound, "Can't find URL: "+tinyUrl.Val)
	}

	logger.MainLogger.Info("Found " + tinyUrl.Val + " as " + fullUrl)

	return &server_proto.FullUrl{Val: fullUrl}, nil
}
