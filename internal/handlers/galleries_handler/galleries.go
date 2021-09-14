package gallerieshandler

import (
	"tung.gallery/internal/services/galleries"
)

type galleryHandler struct {
	Services galleries.GalleriesServiceInterface
}

func NewGalleryHandler(s galleries.GalleriesServiceInterface) *galleryHandler {
	return &galleryHandler{Services: s}
}
