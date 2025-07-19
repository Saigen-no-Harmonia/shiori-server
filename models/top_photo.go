package models

type TopPhoto struct {
	Id           string `json:"id"`
	S3ObjectName string `json:"s3ObjectName"`
}

func NewTopPhoto(
	id string,
	s3ObjectName string,
) *TopPhoto {
	p := new(TopPhoto)
	p.Id = id
	p.S3ObjectName = s3ObjectName

	return p
}
