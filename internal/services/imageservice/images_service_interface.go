package imageservice

import (
	"context"
	"tung.gallery/internal/dt/dto/imagedto"
	"tung.gallery/internal/repo"
)

type ImagesServiceInterface interface {
	GetImageByID(ctx context.Context, userID, imageID int64) (*imagedto.GetImageResponse, error)
	GetImageByGalleryID(ctx context.Context, userID int64, galleryID []int64) (*imagedto.GetImageListResponse, error)
	GetImageByUserID(ctx context.Context, currentUser, userID int64) (*imagedto.GetImageListResponse, error)
	CreateImage(ctx context.Context, req *imagedto.ImageUploadRequest, content string) (string, error)
}

type imageService struct {
	Repo repo.GalleryRepository
}

func NewImageService(r repo.GalleryRepository) *imageService {
	return &imageService{
		Repo: r,
	}
}
