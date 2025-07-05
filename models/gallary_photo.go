package models

type GalleryPhoto struct {
	GalleryPhotoId  string `json:"galleryPhotoId"`
	GalleryPhotoUrl string `json:"galleryPhotoUrl"`
}

func NewGalleryPhoto(
	galleryPhotoId string,
	galleryPhotoUrl string,
) *GalleryPhoto {
	p := new(GalleryPhoto)
	p.GalleryPhotoId = galleryPhotoId
	p.GalleryPhotoUrl = galleryPhotoUrl

	return p
}
