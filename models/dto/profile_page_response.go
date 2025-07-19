package dto

import "shiori-server/models"

type ProfilePageResponse struct {
	PresenterProfiles   []models.PresenterProfile   `json:"presenterProfiles"`
	ParticipantProfiles []models.ParticipantProfile `json:"participantProfiles"`
	NekoProfiles        []models.Neko               `json:"nekoProfiles"`
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
