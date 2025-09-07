package service

import (
	"shiori-server/dto"
	"shiori-server/repository"
)

/** インタフェース */
type AccessService interface {
	GetAccessPage() (*dto.AccessPageResource, error)
}

/** サービス構造体 */
type accessService struct {
	accessRepo repository.AccessRepository
}

/** サービス構造体のコンストラクタ */
func NewAccessService(ar repository.AccessRepository) AccessService {
	return &accessService{accessRepo: ar}
}

/**
 * アクセスページ情報を取得
 */
func (s *accessService) GetAccessPage() (*dto.AccessPageResource, error) {
	// DBからアクセスページ情報を取得
	accessInfo, err := s.accessRepo.Select()
	if err != nil {
		return nil, err
	}

	return dto.MapToAccessPageResource(*accessInfo), nil
}
