package models

type Ie struct {
	Id   string `json:"id" example:"IE000"`
	Name string `json:"name" example:"鈴木"`
}

func NewIe(
	id string,
	name string,
) *Ie {
	i := new(Ie)
	i.Id = id
	i.Name = name

	return i
}
