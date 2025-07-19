package models

type TopPhoto struct {
	Id           string `json:"id"`
	S3ObjectName string `json:"s3ObjectName"`
	PhotoUrl     string `json:"photoUrl"`
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
