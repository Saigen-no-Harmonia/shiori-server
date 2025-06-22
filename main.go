package main

import (
	"shiori-server/database"
	"shiori-server/controllers"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	// .envファイルの読み込み
	err := godotenv.Load()
	if err != nil {
		log.Fatal(".envファイルの読み込みに失敗しました")
	}

	// DBの初期化
	database.InitDB()

	r := gin.Default()

	// ルーティング
	r.GET("/", controllers.getTopInformation)

	// サーバ起動
	r.RUN(":8080") //localhost:8080
}