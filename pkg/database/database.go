package database

import (
	"booking-ticket/logger"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type db struct {
	db *sql.DB
}

func NewDatabase(dsn string) (*db, error) {
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening database connection: %v", err)
	}

	if err := conn.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging database: %v", err)
	}
	// log.Println("Connected to database")
	// logger.Log.Info("Connected to database")
	logger.Info("Connected to database")

	return &db{db: conn}, nil
}

func (d *db) GetConnection() *sql.DB {
	return d.db
}
