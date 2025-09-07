package repository

import (
	"database/sql"
	"shiori-server/model"
)

type GalleryPhotoRepository interface {
	SelectByRange(limit int, offset int) ([]model.GalleryPhoto, error)
}

type galleryPhotoRepository struct {
	db *sql.DB
}

func NewGalleryPhotoRepository(db *sql.DB) GalleryPhotoRepository {
	return &galleryPhotoRepository{db: db}
}

func (r *galleryPhotoRepository) SelectByRange(limit int, offset int) ([]model.GalleryPhoto, error) {
	// DBから画像取得
	rows, err := r.db.Query("SELECT id, s3_object_name, display_number "+
		"FROM M_GALLERY_PHOTO "+
		"WHERE delete_flag = 0 "+
		"ORDER BY display_number ASC "+
		"LIMIT ? OFFSET ?",
		limit,
		offset,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// 取得したデータをモデルに格納
	galleryPhotos := []model.GalleryPhoto{}
	for rows.Next() {
		var p model.GalleryPhoto
		if err := rows.Scan(
			&p.Id,
			&p.S3ObjectName,
			&p.DisplayNumber,
		); err != nil {
			return nil, err
		}

		galleryPhotos = append(galleryPhotos, p)
	}

	return galleryPhotos, nil
}
