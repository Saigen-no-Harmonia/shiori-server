package dto

import "shiori-server/models"

type GalleryPhotoPageResponse struct {
	GalleryPhotos []models.GalleryPhoto `json:"galleryPhotos"`
}

func NewGalleryPhotoPageresponse(
	galleryPhotos []models.GalleryPhoto,
) *GalleryPhotoPageResponse {
	r := new(GalleryPhotoPageResponse)
	r.GalleryPhotos = galleryPhotos

	return r
}
