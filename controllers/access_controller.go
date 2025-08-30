package controllers

import (
	"database/sql"
	"net/http"
	"shiori-server/database"
	"shiori-server/models/resource"
	"shiori-server/util"

	"github.com/gin-gonic/gin"
)

// GetAccessPage アクセスページ情報の取得
// @Summary アクセスページ取得
// @Description アクセス情報を返却。
// @Tags access
// @Accept json
// @Produce json
// @Success 200 {object} resource.AccessPageResource
// @Failure 500 {object} map[string]string
// @Router /access [get]
func GetAccessPage(c *gin.Context) {
	/** トップ画像をDBから取得（レコードは一件のみ） */
	accessRow := database.DB.QueryRow(
		"SELECT venue_id, venue_name, venue_address, venue_access_page_url, latitude, longitude, gathering_spot, gathering_date, starting_date, restaurant_name, restaurant_url " +
			"FROM M_ACCESS_INFO " +
			"LIMIT 1",
	)

	/** 取得結果をDTOにマッピング */
	accessInfo, accessInfoErr := util.MapToAccessInfo(accessRow)
	if accessInfoErr != nil {
		/** 取得できなかった場合 */
		if accessInfoErr == sql.ErrNoRows {
			c.JSON(
				http.StatusNotFound,
				gin.H{"アクセス情報の取得に失敗しました。": accessInfoErr.Error()},
			)
			return
		}
		//** その他のエラーの場合 */
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"アクセス情報取得時に予期せぬエラーが生じました。": accessInfoErr.Error()},
		)
		return
	}

	// ResourceにMap
	resource := resource.MapToAccessPageResource(accessInfo)

	c.JSON(http.StatusOK, resource)
}
