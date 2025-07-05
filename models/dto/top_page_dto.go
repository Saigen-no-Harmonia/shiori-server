package dto

import "shiori-server/models"

type TopPageResponse struct {
	TopPhoto models.TopPhoto   `json:"TopPhoto"`
	Greeting []models.Greeting `json:"Greetings"`
}

func NewTopPageResponse(
	topPhoto models.TopPhoto,
	greeting []models.Greeting,
) *TopPageResponse {
	r := new(TopPageResponse)
	r.TopPhoto = topPhoto
	r.Greeting = greeting

	return r
}
