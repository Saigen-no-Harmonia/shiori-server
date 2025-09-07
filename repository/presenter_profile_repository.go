package repository

import (
	"database/sql"
	"shiori-server/model"
	"shiori-server/util"
)

type PresenterProfileRepository interface {
	SelectAll() ([]model.PresenterProfile, error)
}

type presenterProfileRepository struct {
	db *sql.DB
}

func NewPresenterProfileRepository(db *sql.DB) PresenterProfileRepository {
	return &presenterProfileRepository{db: db}
}

func (r *presenterProfileRepository) SelectAll() ([]model.PresenterProfile, error) {
	/** 主催者プロフィールを取得 */
	presenter, err := r.db.Query(
		"SELECT kaede_flg, ie_id, i.name AS ie_name, photo_s3_object_name, last_name, first_name, last_name_kana, first_name_kana, birth_date, birth_place, job, hobby, ramen, nickname, like_by " +
			"FROM M_PRESENTER_PROFILE p " +
			"INNER JOIN M_IE i " +
			"ON p.ie_id = i.id " +
			"ORDER BY kaede_flg ASC ",
	)
	if err != nil {
		/** データ取得に失敗した場合のエラー */
		if err == sql.ErrNoRows {
			return nil, err
		}
	}
	/** 処理終了後にクローズ */
	defer presenter.Close()

	// presenterRowsをループして、情報を記録
	var presenterProfiles []model.PresenterProfile
	for presenter.Next() {
		p := util.MapToPresenterProfile(presenter)

		presenterProfiles = append(presenterProfiles, p)
	}

	return presenterProfiles, nil
}
