package models

type Neko struct {
	nekoId string `json:"nekoId"`
	ieId string `json:"ieId"`
	photoUrl string `json:"photoUrl"`
	name string `json:"name"`
	age int `json:"age"`
	temperament string `json:"temperament"`
	likeFood string `json:"likeFood"`
}