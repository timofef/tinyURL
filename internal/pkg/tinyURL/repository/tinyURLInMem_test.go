package repository

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTinyUrlInMemoryRepository_Add(t *testing.T) {
	type test struct {
		input         []string
		inserted      string
		expectedError error
	}

	tests := []test{
		{input: []string{"fullUrl", "tinyUrl"}, inserted: "fullUrl", expectedError: nil},
	}

	for _, testCase := range tests {
		repo := TinyUrlInMemoryRepository{DB: make(map[string]string)}
		got := repo.Add(testCase.input[0], testCase.input[1])
		assert.Equal(t, testCase.expectedError, got)

		res := repo.DB[testCase.input[1]]
		assert.Equal(t, testCase.inserted, res)
	}
}

func TestTinyUrlInMemoryRepository_Get(t *testing.T) {
	type test struct {
		input         string
		db            map[string]string
		expectedUrl   string
		expectedError error
	}

	tests := []test{
		{input: "tinyUrl", db: map[string]string{"tinyUrl": "fullUrl"}, expectedUrl: "fullUrl", expectedError: nil},
		{input: "tinyUrl", db: make(map[string]string), expectedUrl: "", expectedError: nil},
	}

	for _, testCase := range tests {
		repo := TinyUrlInMemoryRepository{DB: testCase.db}
		got, err := repo.Get(testCase.input)
		assert.Equal(t, testCase.expectedUrl, got)
		assert.Equal(t, testCase.expectedError, err)
	}
}

func TestTinyUrlInMemoryRepository_CheckIfTinyUrlExists(t *testing.T) {
	type test struct {
		input         string
		db            map[string]string
		expected      bool
		expectedError error
	}

	tests := []test{
		{input: "tinyUrl", db: map[string]string{"tinyUrl": "fullUrl"}, expected: true, expectedError: nil},
		{input: "notTinyUrl", db: map[string]string{"tinyUrl": "fullUrl"}, expected: false, expectedError: nil},
	}

	for _, testCase := range tests {
		repo := TinyUrlInMemoryRepository{DB: testCase.db}
		got, err := repo.CheckIfTinyUrlExists(testCase.input)
		assert.Equal(t, testCase.expected, got)
		assert.Equal(t, testCase.expectedError, err)
	}
}

func TestTinyUrlInMemoryRepository_CheckIfFullUrlExists(t *testing.T) {
	type test struct {
		input           string
		db              map[string]string
		expectedTinyUrl string
		expectedError   error
	}

	tests := []test{
		{input: "fullUrl", db: map[string]string{"tinyUrl": "fullUrl"}, expectedTinyUrl: "tinyUrl", expectedError: nil},
		{input: "notFullUrl", db: map[string]string{"tinyUrl": "fullUrl"}, expectedTinyUrl: "", expectedError: nil},
	}

	for _, testCase := range tests {
		repo := TinyUrlInMemoryRepository{DB: testCase.db}
		got, err := repo.CheckIfFullUrlExists(testCase.input)
		assert.Equal(t, testCase.expectedTinyUrl, got)
		assert.Equal(t, testCase.expectedError, err)
	}
}
