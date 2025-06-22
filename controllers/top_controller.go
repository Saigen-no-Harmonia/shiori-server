package controllers

import (
	"net/http"
	"shiori-server/database"
	"shiori-server/models"
)

func getTopInformation(c *gin.Context) {
	// DBからデータ取得
	rows, err := database.DB.Query(
		"SELECT * FROM M_TOP_PHOTO"
	)

	//取得失敗した場合
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// 呼び出し終了時にDBをクローズ
	defer rows.Close()

	// TOP画像情報を変数に格納
	var topPhoto models.topPhoto
	if err := rows.Scan(&topPhoto.topPhotoId, &topPhoto.topPhotoUrl); if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, topPhoto)

	// TODO これだけだと挨拶返せない。DTOが必要
}
	