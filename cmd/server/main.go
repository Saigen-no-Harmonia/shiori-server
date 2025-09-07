// @title Shiori API
// @version 1.0
// @description 顔合わせアプリのエンドポイントを提供
// @license.name yagi
// @BasePath /
package main

import (
	"log"
	"shiori-server/database"
	"shiori-server/handler"
	"shiori-server/repository"
	"shiori-server/service"
	"shiori-server/util"

	_ "shiori-server/docs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	log.Println("===== Starting Gin Server =====")

	// .envファイルの読み込み
	log.Println("===== Starting load dotenv =====")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(".envファイルの読み込みに失敗しました")
	}

	// S3接続のための初期化処理
	log.Println("===== Starting init s3 =====")
	util.InitS3()

	// DBの初期化
	log.Println("===== Starting init DB =====")
	db := database.InitDB()

	// 依存関係の整理
	topPhotoRepo := repository.NewTopPhotoRepository(db)
	greetingRepo := repository.NewGreetingRepository(db)
	presenterProfileRepo := repository.NewPresenterProfileRepository(db)
	participantProfileRepo := repository.NewParticipantProfileRepository(db)
	nekoProfileRepo := repository.NewNekoProfileRepository(db)
	galleryPhotoRepo := repository.NewGalleryPhotoRepository(db)
	accessRepo := repository.NewAccessRepository(db)

	topService := service.NewTopService(topPhotoRepo, greetingRepo)
	familiesService := service.NewFamiliesService(presenterProfileRepo, participantProfileRepo, &nekoProfileRepo)
	galleryService := service.NewGalleryService(galleryPhotoRepo)
	accessService := service.NewAccessService(accessRepo)

	topHandler := handler.NewTopHandler(topService)
	familiesHandler := handler.NewFamiliesHandler(familiesService)
	galleryHandler := handler.NewGalleryHandler(galleryService)
	accessHandler := handler.NewAccessHandler(accessService)

	// Ginルータを生成
	r := gin.Default()
	r.Use(util.AuthByTokenMiddleware())

	// Swagger UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// ルーティング
	r.GET("/top", topHandler.GetTopPage)
	r.GET("/families", familiesHandler.GetFamilies)
	r.GET("/access", accessHandler.GetAccessPage)
	r.GET("/gallery", galleryHandler.GetGalleryPhotos)

	// サーバ起動
	r.Run(":8080") //localhost:8080
}
