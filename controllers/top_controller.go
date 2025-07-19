package controllers

import (
	"database/sql"
	"net/http"
	"shiori-server/database"
	"shiori-server/models"
	"shiori-server/models/resource"
	"shiori-server/util"

	"github.com/gin-gonic/gin"
)

// GetTopPage トップページ情報の取得
// @Summary トップページ取得
// @Description トップ画像と挨拶文を取得。挨拶文は念のため配列だが、要素は１つしか返さない想定。
// @Tags top
// @Accept json
// @Produce json
// @Success 200 {object} resource.TopPageResource
// @Failure 500 {object} map[string]string
// @Router / [get]
func GetTopPage(c *gin.Context) {
	/** トップ画像をDBから取得（レコードは一件のみ） */
	photoRow := database.DB.QueryRow(
		"SELECT id, s3_object_name " +
			"FROM M_TOP_PHOTO " +
			"WHERE delete_flag = 0 " +
			"ORDER BY id ASC " +
			"LIMIT 1",
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

	// 写真アクセス用URLを取得・格納
	topPhoto.PhotoUrl = util.GetS3AccessUrl("トップ画像")

	// ResourceにMap
	topPhotoResource := resource.MapToTopPhotoResource(topPhoto)

	// 挨拶文をDBから取得
	greetingRow := database.DB.QueryRow(
		"SELECT id, content " +
			"FROM M_GREETING " +
			"WHERE delete_flag = 0 " +
			"ORDER BY id ASC " +
			"LIMIT 1",
	)

	/** 取得結果をDTOにマッピング */
	var greeting models.Greeting
	if greetingErr := greetingRow.Scan(
		&greeting.Id,
		&greeting.Content,
	); greetingErr != nil {
		/** 取得件数が０件の場合のエラー*/
		if greetingErr == sql.ErrNoRows {
			c.JSON(
				http.StatusNotFound,
				gin.H{"error": greetingErr.Error()},
			)
			return
		}
		//** その他のエラーの場合 */
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": greetingErr.Error()},
		)
		return
	}

	// ResourceにMap
	greetingResource := resource.MapToGreetingResource(greeting)

	// Resourceに詰めて返却
	response := resource.NewTopPageResource(*topPhotoResource, *greetingResource)
	c.JSON(http.StatusOK, response)
}
