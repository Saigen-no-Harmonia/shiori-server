package resource

import (
	"shiori-server/models"
	"time"
)

type ParticipantProfileResource struct {
	Id            string    `json:"id" example:"PP001"`
	DisplayNumber int       `json:"displayNumber" example:"2"`
	PhotoUrl      string    `json:"photoUrl" example:"https://tanakapapagazou"`
	LastName      string    `json:"lastName" example:"太郎"`
	FirstName     string    `json:"firstName" example:"田中"`
	LastNameKana  string    `json:"lastNameKana" example:"たろう"`
	FirstNameKana string    `json:"firstNameKana" example:"たなか"`
	BirthDate     time.Time `json:"birthDate" example:"1990-05-22"`
	BirthPlace    string    `json:"birthPlace" example:"北海道"`
	Job           string    `json:"job" example:"ユーチューバー"`
	Hobby         string    `json:"hobby" example:"youtube"`
	Relation      string    `json:"relation" example:"父"`
	LikeFood      string    `json:"likeFood" example:"りんご"`
	Message       string    `json:"message" example:"よろしく！"`
}

func NewParticipantProfileResource(
	id string,
	displayNumber int,
	photoUrl string,
	lastName string,
	firstName string,
	lastNameKana string,
	firstNameKana string,
	birthDate time.Time,
	birthPlace string,
	hobby string,
	job string,
	relation string,
	likeFood string,
	message string,
) *ParticipantProfileResource {
	p := new(ParticipantProfileResource)
	p.Id = id
	p.DisplayNumber = displayNumber
	p.PhotoUrl = photoUrl
	p.LastName = lastName
	p.FirstName = firstName
	p.LastNameKana = lastNameKana
	p.FirstNameKana = firstNameKana
	p.BirthDate = birthDate
	p.BirthPlace = birthPlace
	p.Hobby = hobby
	p.Job = job
	p.Relation = relation
	p.LikeFood = likeFood
	p.Message = message

	return p
}

// ModelをResourceにMap
func MapToParticipantProfileResource(
	p models.ParticipantProfile,
) *ParticipantProfileResource {
	return NewParticipantProfileResource(
		p.Id,
		p.DisplayNumber,
		p.PhotoUrl,
		p.LastName,
		p.FirstName,
		p.LastNameKana,
		p.FirstNameKana,
		p.BirthDate,
		p.BirthPlace,
		p.Hobby,
		p.Job,
		p.Relation,
		p.LikeFood,
		p.Message,
	)
}
