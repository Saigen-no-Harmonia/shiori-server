package dto

import "shiori-server/model"

/** 家族単位で情報をまとめるリソース */
type FamilyResource struct {
	IeId                        string                       `json:"ieId" example:"IE000"`
	IeName                      string                       `json:"ieName" example:"田中"`
	PresenterProfileResource    PresenterProfileResource     `json:"presenterProfile"`
	ParticipantProfileResources []ParticipantProfileResource `json:"participantProfiles"`
	NekoProfileResources        []NekoProfileResource        `json:"nekoProfiles"`
}

func NewFamilyProfileResource(
	ieId string,
	ieName string,
	presenterProfile PresenterProfileResource,
	participantProfiles []ParticipantProfileResource,
	nekoProfiles []NekoProfileResource,
) *FamilyResource {
	f := new(FamilyResource)
	f.IeId = ieId
	f.IeName = ieName
	f.PresenterProfileResource = presenterProfile
	f.ParticipantProfileResources = participantProfiles
	f.NekoProfileResources = nekoProfiles
	return f
}

/** 第１引数に入れた主催者と、その家族のプロフィールを抽出し、FamilyProfileResourceにマッピング */
func CreateFamilyProfileResource(
	presenterProfile model.PresenterProfile,
	participantProfiles []model.ParticipantProfile,
	nekoProfiles []model.Neko,
) *FamilyResource {
	f := new(FamilyResource)

	// 家情報は、主催者を参照
	f.IeId = presenterProfile.IeId
	f.IeName = presenterProfile.IeName

	// 主催者プロフィールをResourceに詰めかえる
	presenterProfileResource := *MapToPresenterPhotoResource(presenterProfile)

	// 参加者プロフィールをResourceに詰めかえる
	var participantProfileResources []ParticipantProfileResource
	for index := 0; index < len(participantProfiles); index++ {
		src := participantProfiles[index]
		dist := MapToParticipantProfileResource(src)
		participantProfileResources = append(participantProfileResources, *dist)
	}

	// ねこプロフィールをResourceに詰めかえる
	var nekoProfileResources []NekoProfileResource
	for index := 0; index < len(nekoProfiles); index++ {
		src := nekoProfiles[index]
		dist := MapToNekoProfileResource(src)
		nekoProfileResources = append(nekoProfileResources, *dist)
	}

	f.PresenterProfileResource = presenterProfileResource
	f.ParticipantProfileResources = participantProfileResources
	f.NekoProfileResources = nekoProfileResources
	return f
}
