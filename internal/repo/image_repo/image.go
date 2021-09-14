package image_repo

import (
	"context"
	"gorm.io/gorm"
	"tung.gallery/internal/dt/entity"
)

type imageRepository struct {
	db *gorm.DB
}

func NewImageRepository(db *gorm.DB) *imageRepository {
	return &imageRepository{
		db,
	}
}

func (r *imageRepository) GetImageByID (ctx context.Context, id uint) (*entity.Images, error) {
	image := &entity.Images{}
	err := r.db.Where("id = ?", id).First(image).Error
	if err != nil {
		return nil, err
	}
	return image, nil
}
