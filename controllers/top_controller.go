package controllers

import (
	"database/sql"
	"net/http"
	"shiori-server/database"
	"shiori-server/models/dto"

	"github.com/gin-gonic/gin"
)

func GetTopPage(c *gin.Context) {
	var topPageResponse dto.TopPageResponse

	/** トップ画像をDBから取得（レコードは一件のみ） */
	photoRow := database.DB.QueryRow(
		"SELECT * " +
			"FROM M_TOP_PHOTO" +
			"WHERE delete_flag = 0" +
			"ORDER BY top_photo_id ASC" +
			"LIMIT ONE",
	)

	/** 取得結果をDTOにマッピング */
	if photoErr := photoRow.Scan(
		&topPageResponse.TopPhoto.TopPhotoId,  // トップ画像ID
		&topPageResponse.TopPhoto.TopPhotoUrl, // トップ画像URL
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
	greetingRow := database.DB.QueryRow(
		"SELECT *" +
			"FROM M_GREETING" +
			"WHERE delete_flag = 0" +
			"ORDER BY greeting_id ASC" +
			"LIMIT ONE",
	)
	// 挨拶文テーブルをクローズ?

	/** 取得結果をDTOにマッピング*/
	if greetingErr := greetingRow.Scan(
		&topPageResponse.Greeting.GreetingId,    // 挨拶文ID
		&topPageResponse.Greeting.DisplayNumber, // 表示順（１のみ）
		&topPageResponse.Greeting.Content,       // 挨拶文
	); greetingErr != nil {
		/** 取得結果が0件の場合のエラー */
		if greetingErr == sql.ErrNoRows {
			c.JSON(
				http.StatusNotFound,
				gin.H{"error": greetingErr.Error()},
			)
		}

		/** その他の場合のエラー */
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": greetingErr.Error()},
		)
	}

	// DTOにつめて返却
	c.JSON(http.StatusOK, topPageResponse)
}
