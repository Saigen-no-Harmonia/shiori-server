package resource

type TopPageResource struct {
	TopPhotoResource TopPhotoResource `json:"topPhoto"`
	GreetingResource GreetingResource `json:"greeting"`
}

func NewTopPageResource(
	topPhotoResource TopPhotoResource,
	greetingResource GreetingResource,
) *TopPageResource {
	r := new(TopPageResource)
	r.TopPhotoResource = topPhotoResource
	r.GreetingResource = greetingResource

	return r
}
