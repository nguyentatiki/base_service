package database

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func Open() *sqlx.DB {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	conn := os.Getenv("MYSQL_CONNECTION")
	db, err := sqlx.Connect("mysql", conn)

	if err != nil {
		fmt.Printf("Error Opening DB: %v \n", err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Printf("Error Pinging DB: %v \n", err)
	}

	fmt.Println("Connected to db!")

	return db
}
