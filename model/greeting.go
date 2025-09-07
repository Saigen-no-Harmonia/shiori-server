package model

type Greeting struct {
	Id            string `json:"id"  example:"GT001"`
	DisplayNumber int    `json:"displayNumber" example:"1"`
	Content       string `json:"content" example:"こんにちは。たなかです。お忙しい中と存じますが云々。"`
}

func NewGreeting(
	id string,
	displayNumber int,
	content string,
) *Greeting {
	g := new(Greeting)
	g.Id = id
	g.DisplayNumber = displayNumber
	g.Content = content

	return g
}
