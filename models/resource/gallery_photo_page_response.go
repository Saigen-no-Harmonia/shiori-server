package resource

type GalleryPhotoPageResource struct {
	GalleryPhotoResources []GalleryPhotoResource `json:"galleryPhotos"`
}

func NewGalleryPhotoPageResource(
	galleryPhotoResources []GalleryPhotoResource,
) *GalleryPhotoPageResource {
	r := new(GalleryPhotoPageResource)
	r.GalleryPhotoResources = galleryPhotoResources

	return r
}
