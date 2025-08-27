package database

import (
	"database/sql"
	"fmt"
	
	_ "github.com/go-sql-driver/mysql"
)

// Connect to the database / *sql.DB is a database connection pool
func Connect() (*sql.DB, error) {
	stringConnection := "golang:golang@tcp(localhost:3306)/learningo?charset=utf8&parseTime=True&loc=Local"

	db, error := sql.Open("mysql", stringConnection)
	if error != nil {
		return nil, error
	}

	if error = db.Ping(); error != nil {
		return nil, fmt.Errorf("database ping error: %v", error)
	}

	return db, nil

}
