package models

type Ie struct {
	IeId   string `json:"ieId"`
	IeName string `json:"ieName"`
}

func NewIe(
	ieId string,
	ieName string,
) *Ie {
	i := new(Ie)
	i.IeId = ieId
	i.IeName = ieName

	return i
}
