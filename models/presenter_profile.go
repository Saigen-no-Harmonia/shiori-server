package models

import "time"

type PresenterProfile struct {
	KaedeFlag         int8      `json:"kaedeFlag"`
	IeId              string    `json:"ieId"`
	IeName            string    `json:"ieName"`
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
	Ramen             string    `json:"ramen"`
	NickName          string    `json:"nickName"`
	LikeBy            string    `json:"likeBy"`
}

func NewPresenterProfile(
	kaedeFlag int8,
	ieId string,
	ieName string,
	photoS3ObjectName string,
	photoUrl string,
	lastName string,
	firstName string,
	lastNameKana string,
	firstNameKana string,
	birthDate time.Time,
	birthPlace string,
	job string,
	hobby string,
	ramen string,
	nickName string,
	likeBy string,
) *PresenterProfile {
	p := new(PresenterProfile)
	p.KaedeFlag = kaedeFlag
	p.IeId = ieId
	p.IeName = ieName
	p.PhotoS3ObjectName = photoS3ObjectName
	p.PhotoUrl = photoUrl
	p.LastName = lastName
	p.FirstName = firstName
	p.LastNameKana = lastNameKana
	p.FirstNameKana = firstNameKana
	p.BirthDate = birthDate
	p.BirthPlace = birthPlace
	p.Job = job
	p.Hobby = hobby
	p.Ramen = ramen
	p.NickName = nickName
	p.LikeBy = likeBy

	return p
}
