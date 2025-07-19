package resource

import "shiori-server/models"

type GalleryPhotoResource struct {
	Id            string `json:"id" example:"GP001"`
	DisplayNumber string `json:"displayNumber" example:"2"`
	PhotoUrl      string `json:"photoUrl" example:"https://s3url"`
}

func NewGalleryPhotoResource(
	id string,
	displayNumber string,
	photoUrl string,
) *GalleryPhotoResource {
	p := new(GalleryPhotoResource)
	p.Id = id
	p.DisplayNumber = displayNumber
	p.PhotoUrl = photoUrl

	return p
}

// ModelをResourceにMap
func MapToGalleryResource(p models.GalleryPhoto) *GalleryPhotoResource {
	return NewGalleryPhotoResource(
		p.Id,
		p.DisplayNumber,
		p.PhotoUrl,
	)
}
