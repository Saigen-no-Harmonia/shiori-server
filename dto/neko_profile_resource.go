package dto

import "shiori-server/model"

type NekoProfileResource struct {
	Id            string `json:"id" example:"NK001"`
	DisplayNumber int    `json:"displayNumber" example:"1"`
	PhotoUrl      string `json:"photoUrl" example:"https://nekoneko"`
	Name          string `json:"name" example:"nekoyagi"`
	Age           int    `json:"age" example:"30"`
	Temperament   string `json:"temperament" example:"それとなくきつい性格"`
	LikeFood      string `json:"likeFood" example:"アーリオ・オーリオ・ペペロンティーノ"`
}

func NewNekoProfileResource(
	id string,
	displayNumber int,
	photoUrl string,
	name string,
	age int,
	temperament string,
	likeFood string,
) *NekoProfileResource {
	n := new(NekoProfileResource)
	n.Id = id
	n.DisplayNumber = displayNumber
	n.PhotoUrl = photoUrl
	n.Name = name
	n.Age = age
	n.Temperament = temperament
	n.LikeFood = likeFood

	return n
}

// ModelをResourceにMap
func MapToNekoProfileResource(
	n model.Neko,
) *NekoProfileResource {
	return NewNekoProfileResource(
		n.Id,
		n.DisplayNumber,
		n.PhotoUrl,
		n.Name,
		n.Age,
		n.Temperament,
		n.LikeFood,
	)
}
