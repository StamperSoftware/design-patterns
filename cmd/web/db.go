package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

const (
	maxOpenDBConn = 25
	maxIdleDBConn = 25
	maxDBLifetime = 5 * time.Minute
)

func initMySQLDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenDBConn)
	db.SetConnMaxLifetime(maxDBLifetime)
	db.SetMaxIdleConns(maxIdleDBConn)

	return db, nil
}
