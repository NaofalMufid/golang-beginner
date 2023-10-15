package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func StartDB() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// host := os.Getenv("HOST")
	// port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	// dbConfig := fmt.Sprintf("%s:%s@/%s?parseTime=true", user, password, dbname)
	dbConfig := fmt.Sprintf("%s:%s@/%s", user, password, dbname)

	db, err := sql.Open("mysql", dbConfig)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// fmt.Println("Successfully connected to database")

	return db, nil
}
