package repo

import (
	"gorm.io/gorm"
	"tung.gallery/internal/dt/entity"
)

type GalleryRepositoryInterface interface {
	CreateGallery(gallery *entity.Galleries) (*entity.Galleries, error)
	ByID(id uint) (*entity.Galleries, error)
	Update(*entity.Galleries) error
	Delete(id uint) error
}

type galleryRepo struct {
	DB *gorm.DB
}

func NewGalleryRepo(db *gorm.DB) GalleryRepositoryInterface {
	return &galleryRepo{DB: db}
}

func (u *galleryRepo) CreateGallery(gallery *entity.Galleries) (*entity.Galleries, error) {
	err := u.DB.Create(gallery).Error
	return gallery, err
}

func (u *galleryRepo) ByID(id uint) (*entity.Galleries, error) {
	gallery := entity.Galleries{}
	gallery.ID = id

	err := u.DB.First(&gallery).Error
	if err != nil {
		return nil, err
	}

	return &gallery, nil
}

func (u *galleryRepo) Update(gallery *entity.Galleries) error {
	err := u.DB.Updates(gallery).Error
	return err
}

func (u *galleryRepo) Delete(id uint) error {
	err := u.DB.Delete(&entity.Galleries{Model: gorm.Model{ID: id}}).Error
	return err
}
