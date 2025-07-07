package controllers

import (
	"database/sql"
	"net/http"
	"shiori-server/database"
	"shiori-server/models"
	"shiori-server/models/dto"

	"github.com/gin-gonic/gin"
)

func GetProfiles(c *gin.Context) {
	/** 主催者プロフィールを取得 */
	presenterRows, presenterErr := database.DB.Query(
		"SELECT *" +
			"FROM M_PRESENTER_PROFILE p" +
			"INNER JOIN M_IE i" +
			"ON p.ie_id = i.ie_id",
	)
	if presenterErr != nil {
		/** データ取得に失敗した場合のエラー */
		if presenterErr == sql.ErrNoRows {
			c.JSON(
				http.StatusNotFound,
				gin.H{"error": presenterErr.Error()},
			)
		}
		/** その他のエラー */
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
			"FROM M_PARTICIPANT_PROFIILE p" +
			"INNER JOIN M_IE i" +
			"ON p.ie_id = i.ie_id" +
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
			"FROM M_NEKO n" +
			"INNER JOIN M_IE i" +
			"ON n.ie_id = i.ie_id" +
			"ORDER BY display_number",
	)
	if nekoErr != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": nekoErr.Error()},
		)
	}
	defer nekoRows.Close()

	// presenterRowsをループして、情報を記録
	var presenterProfiles []models.PresenterProfile
	for presenterRows.Next() {
		var p models.PresenterProfile
		if err := presenterRows.Scan(
			&p.KaedeFlag,
			&p.IeId,
			&p.IeName,
			&p.PresenterPhotoUrl,
			&p.LastName,
			&p.FirstName,
			&p.LastNameKana,
			&p.FirstNameKana,
			&p.BirthDate,
			&p.BirthPlace,
			&p.Job,
			&p.Hobby,
			&p.Ramen,
			&p.NickName,
			&p.LikeBy,
		); err != nil {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{"error": err.Error()},
			)
		}

		presenterProfiles = append(presenterProfiles, p)
	}

	// participantRowsをループして、情報をadd
	var participantProfiles []models.ParticipantProfile
	for participantRows.Next() {
		var p models.ParticipantProfile
		if err := participantRows.Scan(
			&p.ParticipantProfileId,
			&p.IeId,
			&p.IeName,
			&p.DisplayNumber,
			&p.ParticipantPhotoUrl,
			&p.LastName,
			&p.FirstName,
			&p.LastNameKana,
			&p.FirstNameKana,
			&p.BirthDate,
			&p.BirthPlace,
			&p.Job,
			&p.Hobby,
			&p.Relation,
			&p.LikeFood,
			&p.Message,
		); err != nil {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{"error": err.Error()},
			)
		}

		participantProfiles = append(participantProfiles, p)
	}

	// nekoRowsをループして、情報をadd
	var nekoProfiles []models.Neko
	for nekoRows.Next() {
		var p models.Neko
		if err := nekoRows.Scan(
			&p.NekoId,
			&p.IeId,
			&p.IeName,
			&p.DisplayNumber,
			&p.NekoPhotoUrl,
			&p.Name,
			&p.Age,
			&p.Temperament,
			&p.LikeFood,
		); err != nil {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{"error": err.Error()},
			)
		}

		nekoProfiles = append(nekoProfiles, p)
	}

	profilePageResponse := dto.NewProfilePageResponse(presenterProfiles, participantProfiles, nekoProfiles)
	c.JSON(http.StatusOK, profilePageResponse)
}
