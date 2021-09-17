package gallerieshandler

import (
	"tung.gallery/internal/services/galleryservice"
)

type galleryHandler struct {
	Services galleryservice.GalleriesServiceInterface
}

func NewGalleryHandler(s galleryservice.GalleriesServiceInterface) *galleryHandler {
	return &galleryHandler{Services: s}
}
