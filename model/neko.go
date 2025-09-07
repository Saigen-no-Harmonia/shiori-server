package model

type Neko struct {
	Id                string `json:"id" example:"NK001"`
	IeId              string `json:"ieId" example:"IE000"`
	IeName            string `json:"ieName" example:"鈴木"`
	DisplayNumber     int    `json:"displayNumber" example:"1"`
	PhotoS3ObjectName string `json:"photoS3ObjectName" example:"nekogazou.img"`
	PhotoUrl          string `json:"photoUrl" example:"https://nekoneko"`
	Name              string `json:"name" example:"nekoyagi"`
	Age               int    `json:"age" example:"30"`
	Temperament       string `json:"temperament" example:"それとなくきつい性格"`
	LikeFood          string `json:"likeFood" example:"アーリオ・オーリオ・ペペロンティーノ"`
}

func NewNeko(
	id string,
	ieId string,
	ieName string,
	displayNumber int,
	photoS3ObjectName string,
	photoUrl string,
	name string,
	age int,
	temperament string,
	likeFood string,
) *Neko {
	n := new(Neko)
	n.Id = id
	n.IeId = ieId
	n.IeName = ieName
	n.DisplayNumber = displayNumber
	n.PhotoS3ObjectName = photoS3ObjectName
	n.PhotoUrl = photoUrl
	n.Name = name
	n.Age = age
	n.Temperament = temperament
	n.LikeFood = likeFood

	return n
}
