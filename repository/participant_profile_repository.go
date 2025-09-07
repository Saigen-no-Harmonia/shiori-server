package repository

import (
	"database/sql"
	"shiori-server/model"
	"shiori-server/util"
)

type ParticipantProfileRepository interface {
	SelectAll() ([]model.ParticipantProfile, error)
}

type participantProfileRepository struct {
	db *sql.DB
}

func NewParticipantProfileRepository(db *sql.DB) ParticipantProfileRepository {
	return &participantProfileRepository{db: db}
}

func (r *participantProfileRepository) SelectAll() ([]model.ParticipantProfile, error) {

	// 参加者プロフィールを取得
	participants, err := r.db.Query(
		"SELECT p.id, ie_id, i.name AS ie_name, display_number, photo_s3_object_name, last_name, first_name, last_name_kana, first_name_kana, birth_date, birth_place, job, hobby, relation, like_food, message " +
			"FROM M_PARTICIPANT_PROFILE p " +
			"INNER JOIN M_IE i " +
			"ON p.ie_id = i.id " +
			"ORDER BY display_number ASC",
	)
	// データ取得に失敗した場合
	if err != nil {
		return nil, err
	}

	// 処理終了後にクローズ
	defer participants.Close()

	// participantRowsをループして、情報をadd
	var participantProfiles []model.ParticipantProfile
	for participants.Next() {
		p := util.MapToParticipantProfile(participants)
		participantProfiles = append(participantProfiles, p)
	}

	return participantProfiles, nil
}
