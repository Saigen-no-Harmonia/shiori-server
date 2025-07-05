package dto

import "shiori-server/models"

type ProfilePageResponse struct {
	PresenterProfiles   []models.PresenterProfile   `json:"PresenterProfiles"`
	ParticipantProfiles []models.ParticipantProfile `json:"ParticipantProfiles"`
	NekoProfiles        []models.Neko               `json:"NekoProfiles"`
}

func NewProfilePageResponse(
	presenterProfiles []models.PresenterProfile,
	particopantProfiles []models.ParticipantProfile,
	nekoProfiles []models.Neko,
) *ProfilePageResponse {
	r := new(ProfilePageResponse)
	r.PresenterProfiles = presenterProfiles
	r.ParticipantProfiles = particopantProfiles
	r.NekoProfiles = nekoProfiles

	return r
}
