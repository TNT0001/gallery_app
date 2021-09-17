package galleryrepo

import (
	"gorm.io/gorm"
	"tung.gallery/internal/dt/entity"
)

type GalleryRepositoryInterface interface {
	CreateGallery(gallery *entity.Galleries) (*entity.Galleries, error)
	GetGalleryByID(id int64) (*entity.Galleries, error)
	UpdateGallery(*entity.Galleries) error
	DeleteGallery(id int64) error
	GetGalleryByUserID(userID int64) ([]*entity.Galleries, error)
	GetGalleriesByListID(id []int64) ([]*entity.Galleries, error)
}

type galleryRepo struct {
	DB *gorm.DB
}

func NewGalleryRepo(db *gorm.DB) GalleryRepositoryInterface {
	return &galleryRepo{DB: db}
}
