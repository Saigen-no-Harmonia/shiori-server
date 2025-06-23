package dto

import "shiori-server/models"

type TopPageResponse struct {
	TopPhoto models.TopPhoto `json:"topPhoto`
	Greeting models.Greeting `json:"greeting"`
}
