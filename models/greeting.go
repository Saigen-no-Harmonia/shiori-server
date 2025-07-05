package models

type Greeting struct {
	GreetingId    string `json:"greetingId"`
	DisplayNumber int    `json:"displayNumber"`
	Content       string `json:"content"`
}

func NewGreeting(
	greetingId string,
	displayNumber int,
	content string,
) *Greeting {
	g := new(Greeting)
	g.GreetingId = greetingId
	g.DisplayNumber = displayNumber
	g.Content = content

	return g
}
