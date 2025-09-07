package util

import (
	"database/sql"
	"shiori-server/model"
)

/** DBから取得したデータをmodelにマッピングする */

func MapToNekoProfile(nekoRow *sql.Rows) model.Neko {
	var p model.Neko
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
		panic("ねこ情報の処理に失敗しました。")

	}

	return p
}

func MapToParticipantProfile(participantRow *sql.Rows) model.ParticipantProfile {
	var p model.ParticipantProfile
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
		panic("参加者プロフィールの処理に失敗しました。")
	}

	return p
}

func MapToPresenterProfile(presenterRow *sql.Rows) model.PresenterProfile {
	var p model.PresenterProfile
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
		panic("主催者プロフィールの処理に失敗しました。")
	}

	return p
}
