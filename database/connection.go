package database

import (
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
	"os"
)

var DB sql.DB

//DBの初期化処理

func InitDB() {
	//MySQLのログイン情報取得
	dsn := fmt.Sptintf("%s:%s@tcp(%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
	)


	var err error

	//MySQLのログイン認証
	DB, err = sql.Open("mysql", dsn); if err != nil {
		panic(err)
	}

	//MySQLの疎通確認
	if err = DB.Ping(); err != nil {
		panic(err)
	}
}