package main

import (
	"log"
	"shiori-server/controllers"
	"shiori-server/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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
	r.GET("/", controllers.GetTopPage)

	// サーバ起動
	r.RUN(":8080") //localhost:8080
}
