package model

import "time"

type PresenterProfile struct {
	KaedeFlag         int8      `json:"kaedeFlag" example:"0"`
	IeId              string    `json:"ieId" example:"IE000"`
	IeName            string    `json:"ieName" example:"鈴木"`
	PhotoS3ObjectName string    `json:"photoS3ObjectName" example:"suzuki.img"`
	PhotoUrl          string    `json:"photoUrl" example:"https://suzukiphoto"`
	LastName          string    `json:"lastName" example:"花子"`
	FirstName         string    `json:"firstName" example:"鈴木"`
	LastNameKana      string    `json:"lastNameKana" example:"はなこ"`
	FirstNameKana     string    `json:"firstNameKana" example:"すずき"`
	BirthDate         time.Time `json:"birthDate" example:"1999-01-22"`
	BirthPlace        string    `json:"birthPlace" example:"沖縄"`
	Job               string    `json:"job" example:"警官"`
	Hobby             string    `json:"hobby" example:"どろけい"`
	Ramen             string    `json:"ramen" example:"あぶとら"`
	NickName          string    `json:"nickName" example:"たろちゃん"`
	LikeBy            string    `json:"likeBy" example:"全て"`
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
