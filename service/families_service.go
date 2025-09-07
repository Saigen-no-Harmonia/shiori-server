package service

import (
	"shiori-server/dto"
	"shiori-server/model"
	"shiori-server/repository"
	"shiori-server/util"
)

// FamiliesService Familiesページ情報参照サービスのインタフェース
type FamiliesService interface {
	GetFamiliesPage() ([]dto.FamilyResource, error)
}

// familiesService サービス構造体
type familiesService struct {
	presenterProfileRepo   repository.PresenterProfileRepository
	participantProfileRepo repository.ParticipantProfileRepository
	nekoProfileRepo        repository.NekoProfileRepository
}

// NewFamiliesService サービス構造体のコンストラクタ
func NewFamiliesService(
	prp repository.PresenterProfileRepository,
	pap repository.ParticipantProfileRepository,
	np repository.NekoProfileRepository,
) FamiliesService {
	return &familiesService{presenterProfileRepo: prp, participantProfileRepo: pap, nekoProfileRepo: np}
}

// GetFamiliesPage Familiesページ情報を取得
func (s *familiesService) GetFamiliesPage() ([]dto.FamilyResource, error) {
	// 主催者情報を取得
	presenterProfile, err := s.presenterProfileRepo.SelectAll()
	if err != nil {
		return nil, err
	}
	// 写真用の署名入りURLを設定
	for i := range presenterProfile {
		presenterProfile[i].PhotoUrl = util.GetS3AccessUrl(presenterProfile[i].PhotoS3ObjectName)
	}

	// 参加者情報を取得
	participantProfile, err := s.participantProfileRepo.SelectAll()
	if err != nil {
		return nil, err
	}
	// 写真用の署名入りURLを設定
	for i := range participantProfile {
		participantProfile[i].PhotoUrl = util.GetS3AccessUrl(participantProfile[i].PhotoS3ObjectName)
	}

	// ねこ情報を取得
	nekoProfile, err := s.nekoProfileRepo.SelectAll()
	if err != nil {
		return nil, err
	}

	// 写真用の署名入りURLを設定
	for i := range nekoProfile {
		nekoProfile[i].PhotoUrl = util.GetS3AccessUrl(nekoProfile[i].PhotoS3ObjectName)
	}

	// 家族ごとにまとめる
	families := assenbleFamilies(presenterProfile, participantProfile, nekoProfile)

	return families, nil
}

// assenbleFamilies DBから取得した情報を、家族単位でグルーピングする
func assenbleFamilies(
	presenterProfile []model.PresenterProfile,
	participantProfile []model.ParticipantProfile,
	nekoProfile []model.Neko,
) []dto.FamilyResource {

	var families []dto.FamilyResource
	// 主催者ごとに整理する
	for _, p := range presenterProfile {
		// 主催者と同じ家である参加者を抽出
		var ppSlice []model.ParticipantProfile
		for _, pp := range participantProfile {
			if pp.IeId == p.IeId {
				ppSlice = append(ppSlice, pp)
			}
		}
		// 主催者と同じ家であるねこを抽出
		var npSlice []model.Neko
		for _, n := range nekoProfile {
			if n.IeId == p.IeId {
				npSlice = append(npSlice, n)
			}
		}
		families = append(families, *dto.CreateFamilyProfileResource(p, ppSlice, npSlice))
	}

	return families
}
