package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// DBの初期化処理
func InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	log.Println("DB接続開始")
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Println("DBログインでエラーが起きました")
		panic(err)
	}
	if err = DB.Ping(); err != nil {
		log.Println("DB疎通確認でエラーが起きました")
		panic(err)
	}
	log.Println("DB接続成功")
}
