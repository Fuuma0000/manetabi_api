package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// NewDB DB接続
func NewDB() *sql.DB {
	// DB接続
	db, err := sql.Open("mysql", "fuuma:password@tcp(127.0.0.1:3308)/manetabi_db?parseTime=true&loc=Asia%2FTokyo")

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Connected to DB!")
	return db
}

// CloseDB DB切断
func CloseDB(db *sql.DB) {
	db.Close()
	if err := db.Close(); err != nil {
		log.Fatalln(err)
	}
}
