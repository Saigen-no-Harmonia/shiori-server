package resource

import "shiori-server/models"

type GreetingResource struct {
	Content string `json:"content" example:"こんにちは。たなかです。お忙しい中と存じますが云々。"`
}

// コンストラクタ
func NewGreetingResource(
	content string,
) *GreetingResource {
	g := new(GreetingResource)
	g.Content = content

	return g
}

// ModelをResourceにMap
func MapToGreetingResource(g models.Greeting) *GreetingResource {
	return NewGreetingResource(
		g.Content,
	)
}
