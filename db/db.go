package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connection() (*sql.DB, error) {
	if DB != nil {
		err := DB.Ping()
		if err != nil {
			return initDB()
		} else {
			return DB, nil
		}
	} else {
		return initDB()
	}
}

func initDB() (*sql.DB, error) {
	dbUser := "postgres"
	dbPass := ""
	dbHost := "192.168.99.100:5432"
	dbName := "ah_profile"
	conStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", dbUser, dbPass, dbHost, dbName)
	DB, _ = sql.Open("postgres", conStr)
	err := DB.Ping()
	if err != nil {
		return nil, err
	}
	return DB, nil
}
