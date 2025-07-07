package models

type Neko struct {
	NekoId        string `json:"nekoId"`
	IeId          string `json:"ieId"`
	IeName        string `json:"ieName"`
	DisplayNumber int    `json:"displayNumber"`
	NekoPhotoUrl  string `json:"nekoPhotoUrl"`
	Name          string `json:"name"`
	Age           int    `json:"age"`
	Temperament   string `json:"temperament"`
	LikeFood      string `json:"likeFood"`
}

func NewNeko(
	nekoId string,
	ieId string,
	ieName string,
	displayNumber int,
	nekoPhotoUrl string,
	name string,
	age int,
	temperament string,
	likeFood string,
) *Neko {
	n := new(Neko)
	n.NekoId = nekoId
	n.IeId = ieId
	n.IeName = ieName
	n.DisplayNumber = displayNumber
	n.NekoPhotoUrl = nekoPhotoUrl
	n.Name = name
	n.Age = age
	n.Temperament = temperament
	n.LikeFood = likeFood

	return n
}
