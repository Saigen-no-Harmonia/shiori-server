package dto

type GalleryPhotoPageResource struct {
	GalleryPhotoResources []GalleryPhotoResource `json:"galleryPhotos"`
}

func NewGalleryPhotoPageResource(
	galleryPhotoResources []GalleryPhotoResource,
) *GalleryPhotoPageResource {
	if galleryPhotoResources == nil {
		galleryPhotoResources = []GalleryPhotoResource{}
	}
	return &GalleryPhotoPageResource{
		GalleryPhotoResources: galleryPhotoResources,
	}
}
