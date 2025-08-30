package util

import (
	"database/sql"
	"shiori-server/models"
)

func MapToNekoProfile(nekoRow *sql.Rows) models.Neko {
	var p models.Neko
	if err := nekoRow.Scan(
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

		panic("unimplemented")

	}

	return p
}

func MapToParticipantProfile(participantRow *sql.Rows) models.ParticipantProfile {
	var p models.ParticipantProfile
	if err := participantRow.Scan(
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
		panic("参加者情報の処理に失敗しました。")
	}

	return p
}

func MapToPresenterProfile(presenterRow *sql.Rows) models.PresenterProfile {
	var p models.PresenterProfile
	if err := presenterRow.Scan(
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
		panic("主催者情報の処理に失敗しました。")
	}

	return p
}
