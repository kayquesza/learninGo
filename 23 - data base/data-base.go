package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Connection String | "user:password@/name_database"...
	stringConnection := "golang:golang@/learninGo?charset=utf8&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", stringConnection)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("The connection is open.")

	rows, err := db.Query("select * from users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println(rows)

}
