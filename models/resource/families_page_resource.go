package resource

type FamiliesPageResource struct {
	Families []FamilyResource `json:"families"`
}

func NewProfilePageResponse(
	families []FamilyResource,
) *FamiliesPageResource {
	r := new(FamiliesPageResource)
	r.Families = families
	return r
}
