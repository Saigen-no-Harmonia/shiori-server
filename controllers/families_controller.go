package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"shiori-server/database"
	"shiori-server/models"
	"shiori-server/models/resource"
	"shiori-server/util"

	"github.com/gin-gonic/gin"
)

// GetFamilies 家族情報ページの情報を取得
// @Summary 家族情報取得
// @Description 主催者・参加者・猫プロフィール情報を取得
// @Tags families
// @Accept json
// @Produce json
// @Success 200 {object} resource.FamiliesPageResource
// @Failure 500 {object} map[string]string
// @Router /families [get]
func GetFamilies(c *gin.Context) {
	/** 主催者プロフィールを取得 */
	presenterRows, presenterErr := database.DB.Query(
		"SELECT kaede_flg, ie_id, i.name AS ie_name, photo_s3_object_name, last_name, first_name, last_name_kana, first_name_kana, birth_date, birth_place, job, hobby, ramen, nickname, like_by " +
			"FROM M_PRESENTER_PROFILE p " +
			"INNER JOIN M_IE i " +
			"ON p.ie_id = i.id " +
			"ORDER BY kaede_flg ASC ",
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
		"SELECT p.id, ie_id, i.name AS ie_name, display_number, photo_s3_object_name, last_name, first_name, last_name_kana, first_name_kana, birth_date, birth_place, job, hobby, relation, like_food, message " +
			"FROM M_PARTICIPANT_PROFILE p " +
			"INNER JOIN M_IE i " +
			"ON p.ie_id = i.id " +
			"ORDER BY display_number ASC",
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
		"SELECT n.id, ie_id, i.name AS ie_name, display_number, photo_s3_object_name, n.name, age, temperament, like_food " +
			"FROM M_NEKO n " +
			"INNER JOIN M_IE i " +
			"ON n.ie_id = i.id " +
			"ORDER BY display_number ASC ",
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
			&p.PhotoS3ObjectName,
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
		// S3アクセス用URLを取得・格納
		p.PhotoUrl = util.GetS3AccessUrl(p.PhotoS3ObjectName)

		presenterProfiles = append(presenterProfiles, p)
	}

	// participantRowsをループして、情報をadd
	var participantProfiles []models.ParticipantProfile
	for participantRows.Next() {
		var p models.ParticipantProfile
		if err := participantRows.Scan(
			&p.Id,
			&p.IeId,
			&p.IeName,
			&p.DisplayNumber,
			&p.PhotoS3ObjectName,
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

		//s3アクセス用URLを取得・格納
		p.PhotoUrl = util.GetS3AccessUrl(p.PhotoS3ObjectName)

		participantProfiles = append(participantProfiles, p)
	}

	// nekoRowsをループして、情報をadd
	var nekoProfiles []models.Neko
	for nekoRows.Next() {
		var p models.Neko
		if err := nekoRows.Scan(
			&p.Id,
			&p.IeId,
			&p.IeName,
			&p.DisplayNumber,
			&p.PhotoS3ObjectName,
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

		p.PhotoUrl = util.GetS3AccessUrl(p.PhotoS3ObjectName)

		nekoProfiles = append(nekoProfiles, p)
	}

	// レスポンス生成
	var response []resource.FamilyResource

	// DBから取得したデータを、家族単位の情報になるよう整理する
	for i := 0; i < len(presenterProfiles); i++ {
		p := presenterProfiles[i]

		// ParticipantProfileのうち、ieIdが同じものを抽出してResource作成
		var ppSlice []models.ParticipantProfile
		for j := 0; j < len(participantProfiles); j++ {
			if participantProfiles[j].IeId == p.IeId {
				ppSlice = append(ppSlice, participantProfiles[j])
			}
		}

		//NekoProfileのうち、ieIdが同じものを抽出してResource作成
		var npSlice []models.Neko
		for j := 0; j < len(nekoProfiles); j++ {
			if nekoProfiles[j].IeId == p.IeId {
				npSlice = append(npSlice, nekoProfiles[j])
			}
		}

		// FamilyResourceに置き換えてレスポンスに詰める
		response = append(response, *resource.CreateFamilyProfileResource(p, ppSlice, npSlice))
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)

	encoder := json.NewEncoder(c.Writer)
	encoder.SetEscapeHTML(false)

	if err := encoder.Encode(response); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
