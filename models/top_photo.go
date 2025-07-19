package models

type TopPhoto struct {
	Id           string `json:"id" example:"IE000"`
	S3ObjectName string `json:"s3ObjectName" example:"suzuki.img"`
	PhotoUrl     string `json:"photoUrl" example:"https://suzukinogazou"`
}

func NewTopPhoto(
	id string,
	s3ObjectName string,
	photoUrl string,
) *TopPhoto {
	p := new(TopPhoto)
	p.Id = id
	p.S3ObjectName = s3ObjectName
	p.PhotoUrl = photoUrl

	return p
}
