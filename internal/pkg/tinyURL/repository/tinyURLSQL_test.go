package repository

import (
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTinyUrlSqlRepository_Add(t *testing.T) {
	type test struct {
		input         []string
		expectedError error
		db            func() *sql.DB
	}

	tests := []test{
		{
			input:         []string{"fullUrl", "fullUrl"},
			expectedError: errors.New("sql error"),
			db: func() *sql.DB {
				database, mock, err := sqlmock.New()
				if err != nil {
					t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
				}
				mock.ExpectExec("INSERT INTO urls").WithArgs("fullUrl", "fullUrl").WillReturnError(errors.New("sql error"))

				return database
			},
		},
		{
			input:         []string{"fullUrl", "fullUrl"},
			expectedError: nil,
			db: func() *sql.DB {
				database, mock, err := sqlmock.New()
				if err != nil {
					t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
				}
				mock.ExpectExec("INSERT INTO urls").WithArgs("fullUrl", "fullUrl").WillReturnResult(sqlmock.NewResult(1, 1))

				return database
			},
		},
	}

	for _, testCase := range tests {
		repo := TinyUrlSqlRepository{DB: testCase.db()}
		err := repo.Add(testCase.input[0], testCase.input[1])
		assert.Equal(t, testCase.expectedError, err)
	}
}

func TestTinyUrlSqlRepository_Get(t *testing.T) {
	type test struct {
		tinyUrl         string
		expectedFullUrl string
		expectedError   error
		db              func() *sql.DB
	}

	tests := []test{
		{
			tinyUrl:         "fullUrl",
			expectedFullUrl: "",
			expectedError:   errors.New("sql error"),
			db: func() *sql.DB {
				database, mock, err := sqlmock.New()
				if err != nil {
					t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
				}
				mock.ExpectQuery("SELECT fullurl").WithArgs("fullUrl").WillReturnError(errors.New("sql error"))

				return database
			},
		},
		{
			tinyUrl:         "fullUrl",
			expectedFullUrl: "http://google.com/",
			expectedError:   nil,
			db: func() *sql.DB {
				database, mock, err := sqlmock.New()
				if err != nil {
					t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
				}
				rows := sqlmock.NewRows([]string{"fullurl"}).AddRow("http://google.com/")
				mock.ExpectQuery("SELECT fullurl").WithArgs("fullUrl").WillReturnRows(rows)

				return database
			},
		},
	}

	for _, testCase := range tests {
		repo := TinyUrlSqlRepository{DB: testCase.db()}
		got, err := repo.Get(testCase.tinyUrl)
		assert.Equal(t, testCase.expectedFullUrl, got)
		assert.Equal(t, testCase.expectedError, err)
	}
}

func TestTinyUrlSqlRepository_CheckIfTinyUrlExists(t *testing.T) {
	type test struct {
		tinyUrl       string
		expected      bool
		expectedError error
		db            func() *sql.DB
	}

	tests := []test{
		{
			tinyUrl:       "tinyUrl",
			expected:      false,
			expectedError: errors.New("sql error"),
			db: func() *sql.DB {
				database, mock, err := sqlmock.New()
				if err != nil {
					t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
				}
				mock.ExpectQuery("SELECT fullurl").WithArgs("tinyUrl").WillReturnError(errors.New("sql error"))

				return database
			},
		},
		{
			tinyUrl:       "tinyUrl",
			expected:      false,
			expectedError: nil,
			db: func() *sql.DB {
				database, mock, err := sqlmock.New()
				if err != nil {
					t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
				}
				rows := sqlmock.NewRows([]string{"fullurl"})
				mock.ExpectQuery("SELECT fullurl").WithArgs("tinyUrl").WillReturnRows(rows)

				return database
			},
		},
		{
			tinyUrl:       "tinyUrl",
			expected:      true,
			expectedError: nil,
			db: func() *sql.DB {
				database, mock, err := sqlmock.New()
				if err != nil {
					t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
				}
				rows := sqlmock.NewRows([]string{"fullurl"}).AddRow("tinyUrl")
				mock.ExpectQuery("SELECT fullurl").WithArgs("tinyUrl").WillReturnRows(rows)

				return database
			},
		},
	}

	for _, testCase := range tests {
		repo := TinyUrlSqlRepository{DB: testCase.db()}
		got, err := repo.CheckIfTinyUrlExists(testCase.tinyUrl)
		assert.Equal(t, testCase.expected, got)
		assert.Equal(t, testCase.expectedError, err)
	}
}

func TestTinyUrlSqlRepository_CheckIfFullUrlExists(t *testing.T) {
	type test struct {
		fullUrl       string
		expected      string
		expectedError error
		db            func() *sql.DB
	}

	tests := []test{
		{
			fullUrl:       "fullUrl",
			expected:      "",
			expectedError: errors.New("sql error"),
			db: func() *sql.DB {
				database, mock, err := sqlmock.New()
				if err != nil {
					t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
				}
				mock.ExpectQuery("SELECT tinyurl").WithArgs("fullUrl").WillReturnError(errors.New("sql error"))

				return database
			},
		},
		{
			fullUrl:       "fullUrl",
			expected:      "",
			expectedError: nil,
			db: func() *sql.DB {
				database, mock, err := sqlmock.New()
				if err != nil {
					t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
				}
				rows := sqlmock.NewRows([]string{"tinyurl"})
				mock.ExpectQuery("SELECT tinyurl").WithArgs("fullUrl").WillReturnRows(rows)

				return database
			},
		},
		{
			fullUrl:       "fullUrl",
			expected:      "tinyUrl",
			expectedError: nil,
			db: func() *sql.DB {
				database, mock, err := sqlmock.New()
				if err != nil {
					t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
				}
				rows := sqlmock.NewRows([]string{"tinyurl"}).AddRow("tinyUrl")
				mock.ExpectQuery("SELECT tinyurl").WithArgs("fullUrl").WillReturnRows(rows)

				return database
			},
		},
	}

	for _, testCase := range tests {
		repo := TinyUrlSqlRepository{DB: testCase.db()}
		got, err := repo.CheckIfFullUrlExists(testCase.fullUrl)
		assert.Equal(t, testCase.expected, got)
		assert.Equal(t, testCase.expectedError, err)
	}
}
