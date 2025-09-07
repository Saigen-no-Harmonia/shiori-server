package service

import (
	"shiori-server/dto"
	"shiori-server/repository"
	"shiori-server/util"
)

type GalleryService interface {
	GetGalleryPhotos(limit int, offset int) ([]dto.GalleryPhotoResource, error)
}

type galleryService struct {
	galleryRepo repository.GalleryPhotoRepository
}

func NewGalleryService(gp repository.GalleryPhotoRepository) GalleryService {
	return &galleryService{galleryRepo: gp}
}

func (s *galleryService) GetGalleryPhotos(limit int, offset int) ([]dto.GalleryPhotoResource, error) {
	// ギャラリー画像を取得
	galleryPhotos, err := s.galleryRepo.SelectByRange(limit, offset)
	if err != nil {
		return nil, err
	}

	var resource []dto.GalleryPhotoResource
	for i := range galleryPhotos {
		galleryPhotos[i].PhotoUrl = util.GetS3AccessUrl(galleryPhotos[i].S3ObjectName)
		gp := dto.MapToGalleryResource(galleryPhotos[i])
		resource = append(resource, *gp)
	}

	return resource, nil
}
