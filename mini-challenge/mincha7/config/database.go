package config

import (
	"fmt"
	"mincha7/helper"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBInit() *gorm.DB {
	err := godotenv.Load()
	helper.ErrorPanic(err)

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	// dbURI := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
	sqlInfo := fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=True", user, password, host, port, dbName)

	db, err := gorm.Open(mysql.Open(sqlInfo), &gorm.Config{})
	helper.ErrorPanic(err)

	return db
}
