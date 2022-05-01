package repository

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// Table names
const (
	booksTable        = "books"
	authorsTable      = "authors"
	booksAuthorsTable = "books_authors"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

// Initializes DB with given config
func NewMysqlDB(cfg Config) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// Trying to ping DB when connected
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, err
}
