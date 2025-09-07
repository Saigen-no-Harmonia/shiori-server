package repository

import (
	"database/sql"
	"shiori-server/model"
	"shiori-server/util"
)

type NekoProfileRepository interface {
	SelectAll() ([]model.Neko, error)
}

type nekoProfileRepository struct {
	db *sql.DB
}

func NewNekoProfileRepository(db *sql.DB) nekoProfileRepository {
	return nekoProfileRepository{db: db}
}

func (r *nekoProfileRepository) SelectAll() ([]model.Neko, error) {
	// ねこプロフィールを取得
	nekos, err := r.db.Query(
		"SELECT n.id, ie_id, i.name AS ie_name, display_number, photo_s3_object_name, n.name, age, temperament, like_food " +
			"FROM M_NEKO n " +
			"INNER JOIN M_IE i " +
			"ON n.ie_id = i.id " +
			"ORDER BY display_number ASC ",
	)

	// データ取得に失敗した場合
	if err != nil {
		return nil, err
	}
	defer nekos.Close()

	// nekoRowsをループして、情報をadd
	var nekoProfiles []model.Neko
	for nekos.Next() {
		p := util.MapToNekoProfile(nekos)
		nekoProfiles = append(nekoProfiles, p)
	}

	return nekoProfiles, nil
}
