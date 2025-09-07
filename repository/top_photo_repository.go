package repository

import (
	"database/sql"
	"shiori-server/model"
	"shiori-server/util"
)

type TopPhotoRepository interface {
	FindFirst() (*model.TopPhoto, error)
}

type topPhotoRepository struct {
	db *sql.DB
}

func NewTopPhotoRepository(db *sql.DB) TopPhotoRepository {
	return &topPhotoRepository{db: db}
}

func (r *topPhotoRepository) FindFirst() (*model.TopPhoto, error) {
	row := r.db.QueryRow(
		"SELECT id, s3_object_name " +
			"FROM M_TOP_PHOTO " +
			"WHERE delete_flag = 0 " +
			"ORDER BY id ASC " +
			"LIMIT 1",
	)

	var tp model.TopPhoto
	if err := row.Scan(&tp.Id, &tp.S3ObjectName); err != nil {
		return nil, err
	}

	// 署名入りURLを取得
	tp.PhotoUrl = util.GetS3AccessUrl(tp.S3ObjectName)

	return &tp, nil
}
