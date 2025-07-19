package resource

import (
	"shiori-server/models"
	"time"
)

type PresenterProfileResource struct {
	PhotoUrl      string    `json:"photoUrl" example:"https://tarouphoto"`
	LastName      string    `json:"lastName" example:"太郎"`
	FirstName     string    `json:"firstName" example:"田中"`
	LastNameKana  string    `json:"lastNameKana" example:"たろう"`
	FirstNameKana string    `json:"firstNameKana" example:"たなか"`
	BirthDate     time.Time `json:"birthDate" example:"1999-01-22"`
	BirthPlace    string    `json:"birthPlace" example:"沖縄"`
	Job           string    `json:"job" example:"警官"`
	Hobby         string    `json:"hobby" example:"どろけい"`
	Ramen         string    `json:"ramen" example:"あぶとら"`
	NickName      string    `json:"nickName" example:"たろちゃん"`
	LikeBy        string    `json:"likeBy" example:"全て"`
}

// コンストラクタ
func NewPresenterProfileResource(
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
) *PresenterProfileResource {
	p := new(PresenterProfileResource)
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

// ModelをResourceにMap
func MapToPresenterPhotoResource(
	p models.PresenterProfile,
) *PresenterProfileResource {
	return NewPresenterProfileResource(
		p.PhotoUrl,
		p.LastName,
		p.FirstName,
		p.LastNameKana,
		p.FirstNameKana,
		p.BirthDate,
		p.BirthPlace,
		p.Job,
		p.Hobby,
		p.Ramen,
		p.NickName,
		p.LikeBy,
	)
}
