package controllers

import (
	"database/sql"
	"net/http"
	"shiori-server/database"
	"shiori-server/models"
	"shiori-server/models/dto"

	"github.com/gin-gonic/gin"
)

func GetTopPage(c *gin.Context) {
	/** トップ画像をDBから取得（レコードは一件のみ） */
	photoRow := database.DB.QueryRow(
		"SELECT * " +
			"FROM M_TOP_PHOTO" +
			"WHERE delete_flag = 0" +
			"ORDER BY top_photo_id ASC" +
			"LIMIT ONE",
	)

	/** 取得結果をDTOにマッピング */
	var topPhoto models.TopPhoto
	if photoErr := photoRow.Scan(
		&topPhoto.Id,           // トップ画像ID
		&topPhoto.S3ObjectName, // トップ画像URL
	); photoErr != nil {
		/** 取得件数が０件の場合のエラー*/
		if photoErr == sql.ErrNoRows {
			c.JSON(
				http.StatusNotFound,
				gin.H{"error": photoErr.Error()},
			)
			return
		}
		//** その他のエラーの場合 */
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": photoErr.Error()},
		)
		return
	}

	// 挨拶文をDBから取得（レコードは一件のみ）
	greetingRows, greetingErr := database.DB.Query(
		"SELECT *" +
			"FROM M_GREETING" +
			"WHERE delete_flag = 0" +
			"ORDER BY greeting_id ASC",
	)
	if greetingErr != nil {
		/** データ取得できなかった場合のエラー */
		if greetingErr == sql.ErrNoRows {
			c.JSON(
				http.StatusNotFound,
				gin.H{"error": greetingErr.Error()},
			)
			return
		}
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": greetingErr.Error()},
		)
	}
	defer greetingRows.Close()
	// 挨拶文テーブルをクローズ?

	var greetings []models.Greeting
	/** 取得結果を、要素の数だけ、DTOにマッピング*/
	for greetingRows.Next() {
		var greeting models.Greeting
		if err := greetingRows.Scan(
			&greeting.Id,            // 挨拶ID
			&greeting.DisplayNumber, // 表示順
			&greeting.Content,       //挨拶文
		); err != nil {
			c.JSON(http.StatusInternalServerError,
				gin.H{"error": err.Error()})
		}

		/** 配列に格納 */
		greetings = append(greetings, greeting)
	}

	// DTOにつめて返却
	topPageResponse := dto.NewTopPageResponse(topPhoto, greetings)
	c.JSON(http.StatusOK, topPageResponse)
}
