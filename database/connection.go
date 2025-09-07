package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// DBの初期化処理
func InitDB() *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&timeout=5s&readTimeout=5s&writeTimeout=5s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	log.Printf("Connecting to DB with user: %s host: %s port: %s dbname: %s", os.Getenv("DB_USER"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

	log.Println("DB接続開始")
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println("DBログインでエラーが起きました")
		panic(err)
	}
	if err = db.Ping(); err != nil {
		log.Println("DB疎通確認でエラーが起きました")
		panic(err)
	}
	log.Println("DB接続成功")

	return db
}
