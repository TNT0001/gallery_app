package image_repo

import (
	"context"
	"tung.gallery/internal/dt/entity"
)

type ImageReposirotyInterface interface {
	GetImageByID (ctx context.Context, id uint) (*entity.Images, error)
}
