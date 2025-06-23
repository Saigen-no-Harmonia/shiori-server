package controllers

import (
	"net/http"
	"shiori-server/database"
	"shiori-server/models"
	"shiori-server/models/dto"

	"github.com/gin-gonic/gin"
)

func GetTopPage(c *gin.Context) {
	// トップ画像をDBから取得
	photoRow, err := database.DB.Query(
		"SELECT * FROM M_TOP_PHOTO",
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "トップ画像情報へのDBアクセスに失敗しました"})
		return
	}

	// トップ画像テーブルをクローズ
	defer photoRow.Close()

	// 取得データを格納
	var p models.TopPhoto
	if err := photoRow.Scan(&p.TopPhotoId, &p.TopPhotoUrl); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "トップ画像データの格納に失敗しました"})
		return
	}

	// 挨拶文をDBから取得
	greetingRow, err := database.DB.Query(
		"SELECT * FROM M_GREETING",
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "挨拶文へのDBアクセスに失敗しました"})
		return
	}
	// 挨拶文テーブルをクローズ
	defer greetingRow.Close()

	var g models.Greeting
	if err := greetingRow.Scan(&g.GreetingId, &g.DisplayNumber, &g.Content); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "挨拶文データの格納に失敗しました"})
		return
	}

	// DTOにつめて返却
	response := dto.TopPageResponse{
		TopPhoto: p,
		Greeting: g,
	}

	c.JSON(http.StatusOK, response)
}
