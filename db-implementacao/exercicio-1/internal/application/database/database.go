package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var conn *sql.DB

func Connect() error {
	if err := godotenv.Load("../.env.example"); err != nil {
		panic(err)
	}

	// "root:root@tcp(localhost:3306)/my_db"
	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	conn = db

	return nil
}

func GetConnection() *sql.DB {
	return conn
}
