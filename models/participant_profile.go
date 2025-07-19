package models

import "time"

type ParticipantProfile struct {
	Id                string    `json:"id"`
	IeId              string    `json:"ieId"`
	IeName            string    `json:"ieName"`
	DisplayNumber     int       `json:"displayNumber"`
	PhotoS3ObjectName string    `json:"photoS3ObjectName"`
	PhotoUrl          string    `json:"photoUrl"`
	LastName          string    `json:"lastName"`
	FirstName         string    `json:"firstName"`
	LastNameKana      string    `json:"lastNameKana"`
	FirstNameKana     string    `json:"firstNameKana"`
	BirthDate         time.Time `json:"birthDate"`
	BirthPlace        string    `json:"birthPlace"`
	Job               string    `json:"job"`
	Hobby             string    `json:"hobby"`
	Relation          string    `json:"relation"`
	LikeFood          string    `json:"likeFood"`
	Message           string    `json:"message"`
}

func NewParticipantProfile(
	id string,
	ieId string,
	ieName string,
	displayNumber int,
	photoS3ObjectName string,
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
) *ParticipantProfile {
	p := new(ParticipantProfile)
	p.Id = id
	p.IeId = ieId
	p.IeName = ieName
	p.DisplayNumber = displayNumber
	p.PhotoS3ObjectName = photoS3ObjectName
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
