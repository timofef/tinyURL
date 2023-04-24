package delivery

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	mock_delivery "github.com/timofef/tinyURL/internal/pkg/tinyURL/delivery/mocks"
	server "github.com/timofef/tinyURL/internal/pkg/tinyURL/delivery/server"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func TestTinyUrlHandler_Add(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	type test struct {
		fullUrl           *server.FullUrl
		expectedTinyUrl   *server.TinyUrl
		expectedErrorCode codes.Code
		expectedError     string
		usecase           func() *mock_delivery.MockIUsecase
	}

	tests := []test{
		{
			fullUrl:           &server.FullUrl{Val: "not a url"},
			expectedTinyUrl:   nil,
			expectedErrorCode: codes.InvalidArgument,
			expectedError:     "Invalid URL",
			usecase: func() *mock_delivery.MockIUsecase {
				uc := mock_delivery.NewMockIUsecase(mockCtrl)
				return uc
			},
		},
		{
			fullUrl:           &server.FullUrl{Val: "http://google.com/"},
			expectedTinyUrl:   nil,
			expectedErrorCode: codes.Internal,
			expectedError:     "Server error: usecase.Add failed",
			usecase: func() *mock_delivery.MockIUsecase {
				uc := mock_delivery.NewMockIUsecase(mockCtrl)
				uc.EXPECT().
					Add("http://google.com/").
					Return("", errors.New("usecase.Add failed"))
				return uc
			},
		},
		{
			fullUrl:           &server.FullUrl{Val: "http://google.com/"},
			expectedTinyUrl:   &server.TinyUrl{Val: "tiny"},
			expectedErrorCode: codes.OK,
			expectedError:     "",
			usecase: func() *mock_delivery.MockIUsecase {
				uc := mock_delivery.NewMockIUsecase(mockCtrl)
				uc.EXPECT().
					Add("http://google.com/").
					Return("tiny", nil)
				return uc
			},
		},
	}

	for _, testCase := range tests {
		handler := TinyUrlHandler{
			Usecase: testCase.usecase(),
		}
		ctx := context.Background()

		got, err := handler.Add(ctx, testCase.fullUrl)
		assert.Equal(t, status.Error(testCase.expectedErrorCode, testCase.expectedError), err)
		if testCase.expectedTinyUrl != nil {
			assert.Equal(t, testCase.expectedTinyUrl.Val, got.Val)
		} else {
			assert.Equal(t, testCase.expectedTinyUrl, got)
		}
	}
}

func TestTinyUrlHandler_Get(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	type test struct {
		tinyUrl           *server.TinyUrl
		expectedFullUrl   *server.FullUrl
		expectedErrorCode codes.Code
		expectedError     string
		usecase           func() *mock_delivery.MockIUsecase
	}

	tests := []test{
		{
			tinyUrl:           &server.TinyUrl{Val: "not a url"},
			expectedFullUrl:   nil,
			expectedErrorCode: codes.InvalidArgument,
			expectedError:     "Invalid URL",
			usecase: func() *mock_delivery.MockIUsecase {
				uc := mock_delivery.NewMockIUsecase(mockCtrl)
				return uc
			},
		},
		{
			tinyUrl:           &server.TinyUrl{Val: "http://tiny.com/qwerty"},
			expectedFullUrl:   nil,
			expectedErrorCode: codes.Internal,
			expectedError:     "Server error: usecase.Get failed",
			usecase: func() *mock_delivery.MockIUsecase {
				uc := mock_delivery.NewMockIUsecase(mockCtrl)
				uc.EXPECT().
					Get("http://tiny.com/qwerty").
					Return("", errors.New("usecase.Get failed"))
				return uc
			},
		},
		{
			tinyUrl:           &server.TinyUrl{Val: "http://tiny.com/qwerty"},
			expectedFullUrl:   nil,
			expectedErrorCode: codes.NotFound,
			expectedError:     "Can't find URL: http://tiny.com/qwerty",
			usecase: func() *mock_delivery.MockIUsecase {
				uc := mock_delivery.NewMockIUsecase(mockCtrl)
				uc.EXPECT().
					Get("http://tiny.com/qwerty").
					Return("", nil)
				return uc
			},
		},
		{
			tinyUrl:           &server.TinyUrl{Val: "http://tiny.com/qwerty"},
			expectedFullUrl:   &server.FullUrl{Val: "fullUrl"},
			expectedErrorCode: codes.OK,
			expectedError:     "",
			usecase: func() *mock_delivery.MockIUsecase {
				uc := mock_delivery.NewMockIUsecase(mockCtrl)
				uc.EXPECT().
					Get("http://tiny.com/qwerty").
					Return("fullUrl", nil)
				return uc
			},
		},
	}

	for _, testCase := range tests {
		handler := TinyUrlHandler{
			Usecase: testCase.usecase(),
		}
		ctx := context.Background()

		got, err := handler.Get(ctx, testCase.tinyUrl)
		assert.Equal(t, status.Error(testCase.expectedErrorCode, testCase.expectedError), err)
		if testCase.expectedFullUrl != nil {
			assert.Equal(t, testCase.expectedFullUrl.Val, got.Val)
		} else {
			assert.Equal(t, testCase.expectedFullUrl, got)
		}
	}
}
