package repository

import (
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTinyUrlSqlRepository_InitTinyUrlSqlRepository(t *testing.T) {
	type test struct {
		name string
		db   func() *sql.DB
	}

	tests := []test{
		{
			name: "success",
			db: func() *sql.DB {
				database, _, err := sqlmock.New()
				if err != nil {
					t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
				}

				return database
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			db := testCase.db()
			got := InitTinyUrlSqlRepository(db)

			assert.NotNil(t, got)
			assert.Equal(t, db, got.db)
		})
	}
}

func TestTinyUrlSqlRepository_Add(t *testing.T) {
	type test struct {
		name          string
		input         []string
		expectedError error
		db            func() *sql.DB
	}

	tests := []test{
		{
			name:          "sql_error",
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
			name:          "success",
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
		t.Run(testCase.name, func(t *testing.T) {
			repo := TinyUrlSqlRepository{db: testCase.db()}
			err := repo.Add(testCase.input[0], testCase.input[1])

			assert.Equal(t, testCase.expectedError, err)
		})
	}
}

func TestTinyUrlSqlRepository_Get(t *testing.T) {
	type test struct {
		name            string
		tinyUrl         string
		expectedFullUrl string
		expectedError   error
		db              func() *sql.DB
	}

	tests := []test{
		{
			name:            "sql_error",
			tinyUrl:         "tinyUrl",
			expectedFullUrl: "",
			expectedError:   errors.New("sql error"),
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
			name:            "success",
			tinyUrl:         "tinyUrl",
			expectedFullUrl: "http://google.com/",
			expectedError:   nil,
			db: func() *sql.DB {
				database, mock, err := sqlmock.New()
				if err != nil {
					t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
				}
				rows := sqlmock.NewRows([]string{"fullurl"}).AddRow("http://google.com/")
				mock.ExpectQuery("SELECT fullurl").WithArgs("tinyUrl").WillReturnRows(rows)

				return database
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			repo := TinyUrlSqlRepository{db: testCase.db()}
			got, err := repo.Get(testCase.tinyUrl)

			assert.Equal(t, testCase.expectedFullUrl, got)
			assert.Equal(t, testCase.expectedError, err)
		})
	}
}

func TestTinyUrlSqlRepository_CheckIfTinyUrlExists(t *testing.T) {
	type test struct {
		name          string
		tinyUrl       string
		expected      bool
		expectedError error
		db            func() *sql.DB
	}

	tests := []test{
		{
			name:          "sql_error",
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
			name:          "tinyurl_not_exist",
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
			name:          "tinyurl_exist",
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
		t.Run(testCase.name, func(t *testing.T) {
			repo := TinyUrlSqlRepository{db: testCase.db()}
			got, err := repo.CheckIfTinyUrlExists(testCase.tinyUrl)

			assert.Equal(t, testCase.expected, got)
			assert.Equal(t, testCase.expectedError, err)
		})
	}
}

func TestTinyUrlSqlRepository_CheckIfFullUrlExists(t *testing.T) {
	type test struct {
		name          string
		fullUrl       string
		expected      string
		expectedError error
		db            func() *sql.DB
	}

	tests := []test{
		{
			name:          "sql_error",
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
			name:          "fullurl_not_exist",
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
			name:          "fullurl_exist",
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
		t.Run(testCase.name, func(t *testing.T) {
			repo := TinyUrlSqlRepository{db: testCase.db()}
			got, err := repo.CheckIfFullUrlExists(testCase.fullUrl)

			assert.Equal(t, testCase.expected, got)
			assert.Equal(t, testCase.expectedError, err)
		})
	}
}
