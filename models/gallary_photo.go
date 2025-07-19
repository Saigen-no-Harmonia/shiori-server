package models

type GalleryPhoto struct {
	Id           string `json:"id"`
	S3ObjectName string `json:"s3ObjectName"`
	PhotoUrl     string `json:"photoUrl"`
}

func NewGalleryPhoto(
	id string,
	s3ObjectName string,
	photoUrl string,
) *GalleryPhoto {
	p := new(GalleryPhoto)
	p.Id = id
	p.S3ObjectName = s3ObjectName
	p.PhotoUrl = photoUrl

	return p
}
