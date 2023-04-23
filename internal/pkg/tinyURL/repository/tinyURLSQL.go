package repository

import (
	"database/sql"
	_ "github.com/jackc/pgx"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/timofef/tinyURL/internal/tinyURL/logger"

	"time"
)

type TinyUrlSqlRepository struct {
	DB *sql.DB
}

func InitPostgres(dbUrl string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dbUrl)
	if err != nil {
		return nil, err
	}

	wait := 10
	for i := 0; i < wait; i++ {
		logger.MainLogger.Info("Pinging database...")
		if err = db.Ping(); err == nil {
			logger.MainLogger.Info("Success")
			break
		}
		time.Sleep(1 * time.Second)
	}
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (r *TinyUrlSqlRepository) Add(fullUrl, tinyUrl string) error {
	_, err := r.DB.Exec(`INSERT INTO urls (fullurl, tinyurl) VALUES ($1, $2)`, fullUrl, tinyUrl)

	return err
}

func (r *TinyUrlSqlRepository) Get(tinyUrl string) (string, error) {
	rows, err := r.DB.Query(`SELECT fullurl FROM urls WHERE tinyurl = $1`, tinyUrl)
	if err != nil {
		return "", err
	}
	var fullUrl string
	for rows.Next() {
		if err = rows.Scan(&fullUrl); err != nil {
			return "", err
		}
	}

	return fullUrl, nil
}

func (r *TinyUrlSqlRepository) CheckIfFullUrlExists(fullUrl string) (string, error) {
	rows, err := r.DB.Query(`SELECT tinyurl FROM urls WHERE fullurl = $1`, fullUrl)
	if err != nil {
		return "", err
	}

	var tinyUrl string
	for rows.Next() {
		err = rows.Scan(&tinyUrl)
	}
	if err != nil {
		return "", err
	}
	if len(tinyUrl) == 0 {
		return "", nil
	}

	return tinyUrl, nil
}

func (r *TinyUrlSqlRepository) CheckIfTinyUrlExists(tinyUrl string) (bool, error) {
	rows, err := r.DB.Query(`SELECT tinyUrl FROM urls WHERE tinyurl = $1`, tinyUrl)
	if err != nil {
		return false, err
	}

	var fullUrl string
	for rows.Next() {
		err = rows.Scan(&fullUrl)
	}
	if err != nil {
		return false, err
	}
	if len(fullUrl) == 0 {
		return false, nil
	}

	return true, nil
}
