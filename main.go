// @title Shiori API
// @version 1.0
// @description 顔合わせアプリのエンドポイントを提供
// @license.name yagi
// @BasePath /
package main

import (
	"log"

	_ "shiori-server/docs"
)

func main() {
	log.Println("Hello, Gin server starting")
	// log.Println("===== Starting Gin Server =====")

	// // .envファイルの読み込み
	// log.Println("===== Starting load dotenv =====")
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal(".envファイルの読み込みに失敗しました")
	// }

	// // S3接続のための初期化処理
	// log.Println("===== Starting init s3 =====")
	// util.InitS3()

	// // DBの初期化
	// log.Println("===== Starting init DB =====")
	// database.InitDB()

	// r := gin.Default()

	// // Swagger UI
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// // ルーティング
	// r.GET("/", controllers.GetTopPage)
	// r.GET("/families", controllers.GetFamilies)
	// r.GET("/access", controllers.GetAccessPage)
	// r.GET("/gallery", controllers.GetGalleryPhotos)

	// // サーバ起動
	// r.Run(":8080") //localhost:8080
}
