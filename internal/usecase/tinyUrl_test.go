package usecase

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/timofef/tinyURL/internal/usecase/mocks"
	"github.com/timofef/tinyURL/internal/utils"
	"testing"
)

func GenerateMock() string {
	return "0123456789"
}

func TestTinyUrlUsecase_InitTinyUrlUsecase(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	type test struct {
		name          string
		baseUrl       string
		generatorFunc func() string
		repository    func() *mock_usecase.MockIRepository
	}

	tests := []test{
		{
			name:          "success",
			baseUrl:       "http://base.com/",
			generatorFunc: GenerateMock,
			repository: func() *mock_usecase.MockIRepository {
				uc := mock_usecase.NewMockIRepository(mockCtrl)
				return uc
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			mockRepo := testCase.repository()
			got := InitTinyUrlUsecase(testCase.baseUrl, mockRepo, testCase.generatorFunc)

			assert.NotNil(t, got)
			assert.Equal(t, testCase.baseUrl, got.baseUrl)
			assert.Equal(t, mockRepo, got.repository) // Can't compare functions
			assert.NotNil(t, got.generateTinyUrl)
		})
	}
}

func TestTinyUrlUsecase_Add(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	type test struct {
		name            string
		fullUrl         string
		expectedTinyUrl string
		expectedError   error
		baseUrl         string
		repository      func() *mock_usecase.MockIRepository
	}

	tests := []test{
		{
			name:            "CheckIfFullUrlExists_failed",
			fullUrl:         "http://google.com/",
			expectedTinyUrl: "",
			expectedError:   errors.New("failed repository.CheckIfFullUrlExists"),
			baseUrl:         "http://base.com/",
			repository: func() *mock_usecase.MockIRepository {
				uc := mock_usecase.NewMockIRepository(mockCtrl)
				uc.EXPECT().
					CheckIfFullUrlExists("http://google.com/").
					Return("", errors.New("failed repository.CheckIfFullUrlExists"))
				return uc
			},
		},
		{
			name:            "CheckIfTinyUrlExists_failed",
			fullUrl:         "http://google.com/",
			expectedTinyUrl: "",
			expectedError:   errors.New("failed repository.CheckIfTinyUrlExists"),
			baseUrl:         "http://base.com/",
			repository: func() *mock_usecase.MockIRepository {
				uc := mock_usecase.NewMockIRepository(mockCtrl)
				first := uc.EXPECT().
					CheckIfFullUrlExists("http://google.com/").
					Return("", nil)
				second := uc.EXPECT().
					CheckIfTinyUrlExists("0123456789").
					Return(false, errors.New("failed repository.CheckIfTinyUrlExists"))
				gomock.InOrder(first, second)
				return uc
			},
		},
		{
			name:            "Add_failed",
			fullUrl:         "http://google.com/",
			expectedTinyUrl: "",
			expectedError:   errors.New("failed repository.Add"),
			baseUrl:         "http://base.com/",
			repository: func() *mock_usecase.MockIRepository {
				uc := mock_usecase.NewMockIRepository(mockCtrl)
				first := uc.EXPECT().
					CheckIfFullUrlExists("http://google.com/").
					Return("", nil)
				second := uc.EXPECT().
					CheckIfTinyUrlExists("0123456789").
					Return(false, nil)
				third := uc.EXPECT().
					Add("http://google.com/", "0123456789").
					Return(errors.New("failed repository.Add"))
				gomock.InOrder(first, second, third)
				return uc
			},
		},
		{
			name:            "tinyurl_already_existed",
			fullUrl:         "http://google.com/",
			expectedTinyUrl: "http://base.com/0123456789",
			expectedError:   nil,
			baseUrl:         "http://base.com/",
			repository: func() *mock_usecase.MockIRepository {
				uc := mock_usecase.NewMockIRepository(mockCtrl)
				uc.EXPECT().
					CheckIfFullUrlExists("http://google.com/").
					Return("0123456789", nil)
				return uc
			},
		},
		{
			name:            "successfully_created_tinyurl",
			fullUrl:         "http://google.com/",
			expectedTinyUrl: "http://base.com/0123456789",
			expectedError:   nil,
			baseUrl:         "http://base.com/",
			repository: func() *mock_usecase.MockIRepository {
				uc := mock_usecase.NewMockIRepository(mockCtrl)
				first := uc.EXPECT().
					CheckIfFullUrlExists("http://google.com/").
					Return("", nil)
				second := uc.EXPECT().
					CheckIfTinyUrlExists("0123456789").
					Return(false, nil)
				third := uc.EXPECT().
					Add("http://google.com/", "0123456789").
					Return(nil)
				gomock.InOrder(first, second, third)
				return uc
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			usecase := TinyUrlUsecase{
				baseUrl:         testCase.baseUrl,
				repository:      testCase.repository(),
				generateTinyUrl: GenerateMock,
			}

			got, err := usecase.Add(testCase.fullUrl)

			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedTinyUrl, got)
		})
	}
}

func TestTinyUrlUsecase_Get(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	type test struct {
		name            string
		tinyUrl         string
		expectedFullUrl string
		expectedError   error
		baseUrl         string
		repository      func() *mock_usecase.MockIRepository
	}

	tests := []test{
		{
			name:            "repository_failed",
			tinyUrl:         "http://base.com/0123456789",
			expectedFullUrl: "",
			expectedError:   errors.New("failed repo"),
			baseUrl:         "http://base.com/",
			repository: func() *mock_usecase.MockIRepository {
				uc := mock_usecase.NewMockIRepository(mockCtrl)
				uc.EXPECT().
					Get("0123456789").
					Return("", errors.New("failed repo"))
				return uc
			},
		},
		{
			name:            "success",
			tinyUrl:         "http://base.com/0123456789",
			expectedFullUrl: "fullUrl",
			expectedError:   nil,
			baseUrl:         "http://base.com/",
			repository: func() *mock_usecase.MockIRepository {
				uc := mock_usecase.NewMockIRepository(mockCtrl)
				uc.EXPECT().
					Get("0123456789").
					Return("fullUrl", nil)
				return uc
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			usecase := TinyUrlUsecase{
				baseUrl:         testCase.baseUrl,
				repository:      testCase.repository(),
				generateTinyUrl: GenerateMock,
			}

			got, err := usecase.Get(testCase.tinyUrl)

			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedFullUrl, got)
		})
	}
}

func TestTinyUrlUsecase_Generator(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		randStr1 := utils.GenerateString()
		randStr2 := utils.GenerateString()

		assert.NotEqual(t, randStr1, randStr2)
	})
}
