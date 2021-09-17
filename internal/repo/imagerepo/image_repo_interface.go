package imagerepo

import (
	"context"
	"gorm.io/gorm"
	"tung.gallery/internal/dt/dto/imagedto"
	"tung.gallery/internal/dt/entity"
)

type ImageReposirotyInterface interface {
	GetImageByID(ctx context.Context, id int64) (*entity.Images, error)
	CreateImage(ctx context.Context, req *imagedto.ImageUploadRequest) error
	GetImagesByGalleryID(ctx context.Context, galleryID []int64) ([]*entity.Images, error)
}

type imageRepository struct {
	db *gorm.DB
}

func NewImageRepository(db *gorm.DB) *imageRepository {
	return &imageRepository{
		db,
	}
}
