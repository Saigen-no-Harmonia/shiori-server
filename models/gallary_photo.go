package models

type GalleryPhoto struct {
	galleryPhotoId  string `json:"galleryPhotoId"`
	galleryPhotoUrl string `json:"galleryPhotoUrl"`
}

func NewGalleryPhoto(
	galleryPhotoId string,
	galleryPhotoUrl string,
) *GalleryPhoto {
	p := new(GalleryPhoto)
	p.galleryPhotoId = galleryPhotoId
	p.galleryPhotoUrl = galleryPhotoUrl

	return p
}
