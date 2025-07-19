package models

type GalleryPhoto struct {
	Id            string `json:"id" example:"GP001"`
	S3ObjectName  string `json:"s3ObjectName" example:"galleryPhoto.img"`
	DisplayNumber string `json:"displayNumber" example:"2"`
	PhotoUrl      string `json:"photoUrl" example:"https://s3url"`
}

func NewGalleryPhoto(
	id string,
	s3ObjectName string,
	displayNumber string,
	photoUrl string,
) *GalleryPhoto {
	p := new(GalleryPhoto)
	p.Id = id
	p.S3ObjectName = s3ObjectName
	p.DisplayNumber = displayNumber
	p.PhotoUrl = photoUrl

	return p
}
