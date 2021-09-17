package galleryservice

import (
	"tung.gallery/internal/dt/dto/gallerydto"
	"tung.gallery/internal/repo/galleryrepo"
)

type GalleriesServiceInterface interface {
	CreateGallery(userID int64, req *gallerydto.GalleryCreateRequest) (*gallerydto.GalleryCreateResponse, error)
	Update(userID int64, req *gallerydto.GalleryUpdateRequest) (*gallerydto.GalleryUpdateResponse, error)
	Delete(userID, galleryID int64) (*gallerydto.GalleryDeleteResponse, error)
	GetGalleryByID(userID, galleryID int64) (*gallerydto.ShowGalleryResponse, error)
	GetAllGalleriesByUserID(currentUserID, userID int64) (*gallerydto.ShowAllGalleryByUserIDResponse, error)
}

type galleryService struct {
	Repo galleryrepo.GalleryRepositoryInterface
}

func NewGalleryService(r galleryrepo.GalleryRepositoryInterface) GalleriesServiceInterface {
	return &galleryService{
		Repo: r,
	}
}
