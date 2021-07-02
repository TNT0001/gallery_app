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
	ByUserID(userID uint) ([]entity.Galleries, error)
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
	gallery := entity.Galleries{}
	gallery.ID = id
	err := u.DB.Unscoped().Delete(&gallery).Error
	return err
}

func (u *galleryRepo) ByUserID(userID uint) ([]entity.Galleries, error) {
	galleries := make([]entity.Galleries, 0)

	err := u.DB.Find(&galleries, "galleries.user_id =?", userID).Error
	if err != nil {
		return []entity.Galleries{}, err
	}

	return galleries, nil
}
