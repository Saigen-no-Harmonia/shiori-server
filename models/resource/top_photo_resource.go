package resource

import "shiori-server/models"

type TopPhotoResource struct {
	PhotoUrl string `json:"photoUrl" example:"https://suzukinogazou"`
}

// コンストラクタ
func NewTopPhotoResource(
	photoUrl string,
) *TopPhotoResource {
	p := new(TopPhotoResource)
	p.PhotoUrl = photoUrl

	return p
}

// ModelをResourceにマップ
func MapToTopPhotoResource(
	p models.TopPhoto,
) *TopPhotoResource {
	return NewTopPhotoResource(
		p.PhotoUrl,
	)
}
