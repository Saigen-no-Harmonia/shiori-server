package models

type ParticipantProfile struct {
	participantProfileId string `json:"participantProfileId"`
	ieId string `json:"ieId"`
	displayNumber int `json:"displayNumber"`
	photoUrl string `json:"photoUrl"`
	lastName string `json:"lastName"`
	firstName string `json:"firstName"`
	lastNameKana string `json:"lastNameKana"`
	firstNameKana string `json:"firstNameKana`
	birthDate time.Time `json:"birthDate"`
	birthPlace string `json:"birthPlace"`
	relation string `json:"relation"`
	job string `json:"job"`
	hobby string `json:"hobby"`
	likeFood string `json:"likeFood"`
	message string `json:"message"`
}