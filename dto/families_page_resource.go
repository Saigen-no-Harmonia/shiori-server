package dto

type FamiliesPageResource struct {
	Families []FamilyResource `json:"families"`
}

func NewFamiliesPageResponse(
	families []FamilyResource,
) *FamiliesPageResource {
	r := new(FamiliesPageResource)
	r.Families = families
	return r
}
