package galleries

import (
	"tung.gallery/internal/dt/dto/gallery_dto"
	"tung.gallery/internal/repo/gallery_repo"
)

type GalleriesServiceInterface interface {
	CreateGallery(userID uint, req *gallery_dto.GalleryCreateRequest) (*gallery_dto.GalleryCreateResponse, error)
	Update(userID uint, req *gallery_dto.GalleryUpdateRequest) (*gallery_dto.GalleryUpdateResponse, error)
	Delete(userID, galleryID uint) (*gallery_dto.GalleryDeleteResponse, error)
	GetGalleryByID(userID, galleryID uint) (*gallery_dto.ShowGalleryResponse, error)
	GetAllGalleriesByUserID(currentUserID, userID uint) (*gallery_dto.ShowAllGalleryByUserIDResponse, error)
}

type galleryService struct {
	Repo gallery_repo.GalleryRepositoryInterface
}

func NewGalleryService(r gallery_repo.GalleryRepositoryInterface) GalleriesServiceInterface {
	return &galleryService{
		Repo: r,
	}
}
