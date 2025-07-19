package models

type Greeting struct {
	Id            string `json:"id"`
	DisplayNumber int    `json:"displayNumber"`
	Content       string `json:"content"`
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
