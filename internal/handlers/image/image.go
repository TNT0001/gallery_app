package image

import (
	"tung.gallery/internal/services/imageservice"
)

type imageHandler struct {
	Services imageservice.ImagesServiceInterface
}

func NewImageHandler(s imageservice.ImagesServiceInterface) *imageHandler {
	return &imageHandler{Services: s}
}
