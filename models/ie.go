package models

type Ie struct {
	Id   string `json:"id"`
	Name string `json:"name"`
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
