package models

type Neko struct {
	Id                string `json:"id"`
	IeId              string `json:"ieId"`
	IeName            string `json:"ieName"`
	DisplayNumber     int    `json:"displayNumber"`
	PhotoS3ObjectName string `json:"photoS3ObjectName"`
	PhotoUrl          string `json:"photoUrl"`
	Name              string `json:"name"`
	Age               int    `json:"age"`
	Temperament       string `json:"temperament"`
	LikeFood          string `json:"likeFood"`
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
