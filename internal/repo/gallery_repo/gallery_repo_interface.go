package gallery_repo

import "tung.gallery/internal/dt/entity"

type GalleryRepositoryInterface interface {
	CreateGallery(gallery *entity.Galleries) (*entity.Galleries, error)
	ByID(id uint) (*entity.Galleries, error)
	Update(*entity.Galleries) error
	Delete(id uint) error
	ByUserID(userID uint) ([]*entity.Galleries, error)
}
