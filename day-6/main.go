package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	user     = "root"
	password = ""
	dbname   = "gochort-day6"
)

var (
	db  *sql.DB
	err error
)

func main() {
	config := fmt.Sprintf("%s:%s@/%s", user, password, dbname)
	db, err = sql.Open("mysql", config)

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	fmt.Println("Successfully conected to database")
}
