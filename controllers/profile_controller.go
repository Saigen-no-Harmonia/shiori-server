package controllers

import (
	"net/http"
	"shiori-server/database"
	"shiori-server/models/dto"

	"github.com/gin-gonic/gin"
)

func GetProfiles(c *gin.Context) {
	/** 主催者プロフィールを取得 */
	presenterRows, presenterErr := database.DB.Query(
		"SELECT * FROM M_PRESENTER_PROFILE",
	)
	if presenterErr != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": presenterErr.Error()},
		)
		return
	}
	/** 処理終了後にクローズ */
	defer presenterRows.Close()

	/** 参加者プロフィールを取得 */
	participantRows, participantErr := database.DB.Query(
		"SELECT *" +
			"FROM M_PARTICIPANT_PROFIILE" +
			"ORDER BY display_number",
	)
	if participantErr != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": participantErr.Error()},
		)
	}
	/** 処理終了後にクローズ */
	defer participantRows.Close()

	nekoRows, nekoErr := database.DB.Query(
		"SELECT *" +
			"FROM M_NEKO" +
			"ORDER BY neko_id",
	)

	if nekoErr != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": nekoErr.Error()},
		)
	}

	var profilePageResponse dto.ProfilePageResponse

	// presenterRowsをループして、情報をadd

	// participantRowsをループして、情報をadd

	// nekoRowsをループして、情報をadd

	return
}
