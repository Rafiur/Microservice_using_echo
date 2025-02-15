package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "abcd1234"
	dbname   = "Employee"
)

func DBconnection() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host = %s port = %d user = %s "+"password = %s dbname = %s sslmode = disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging database: %w", err)
	}

	return db, nil
}