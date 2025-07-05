package models

type TopPhoto struct {
	TopPhotoId  string `json:"topPhotoId"`
	TopPhotoUrl string `json:"topPhotoUrl"`
}

func NewTopPhoto(
	topPhotoId string,
	topPhotoUrl string,
) *TopPhoto {
	p := new(TopPhoto)
	p.TopPhotoId = topPhotoId
	p.TopPhotoUrl = topPhotoUrl

	return p
}
