package service

/**
 * TOPページ情報参照サービス
 */

import (
	"shiori-server/dto"
	"shiori-server/repository"
)

/** インタフェース */
type TopService interface {
	GetTopPage() (*dto.TopPageResource, error)
}

/** サービス構造体 */
type topService struct {
	topPhotoRepo repository.TopPhotoRepository
	greetingRepo repository.GreetingRepository
}

/** サービス構造体のコンストラクタ */
func NewTopService(tp repository.TopPhotoRepository, gr repository.GreetingRepository) TopService {
	return &topService{topPhotoRepo: tp, greetingRepo: gr}
}

/**
 * TOPページ情報を取得
 */
func (s *topService) GetTopPage() (*dto.TopPageResource, error) {
	// DBからトップ画像取得
	photo, err := s.topPhotoRepo.FindFirst()
	if err != nil {
		return nil, err
	}

	// ResourceにMap
	photoResource := dto.MapToTopPhotoResource(*photo)

	greeting, err := s.greetingRepo.FindFirst()
	if err != nil {
		return nil, err
	}

	greetingResource := dto.MapToGreetingResource(*greeting)

	return dto.NewTopPageResource(*photoResource, *greetingResource), nil

}
