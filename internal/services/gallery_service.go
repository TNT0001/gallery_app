package services

import (
	"tung.gallery/internal/dt/dto"
	"tung.gallery/internal/dt/entity"
	"tung.gallery/internal/repo"
)

type GalleriesServiceInterface interface {
	CreateGallery(entity.Users, dto.GalleryCreateRequest) (dto.GalleryCreateResponse, error)
	ShowGallery(id uint) (dto.ShowGalleryResponse, error)
	Update(dto.gall)
}

type galleryService struct {
	Repo repo.GalleryRepositoryInterface
}

func NewGalleryService(r repo.GalleryRepositoryInterface) GalleriesServiceInterface {
	return &galleryService{
		Repo: r,
	}
}

func (g *galleryService) CreateGallery(user entity.Users, req dto.GalleryCreateRequest) (dto.GalleryCreateResponse, error) {
	gallery := entity.Galleries{
		UserID: user.ID,
		Title:  req.Title,
	}

	res, err := g.Repo.CreateGallery(&gallery)
	if err != nil {
		return dto.GalleryCreateResponse{
			Title: req.Title,
			Alert: dto.Alert{
				Level:   AlertLvlInfo,
				Message: err.Error(),
			},
		}, err
	}

	return dto.GalleryCreateResponse{
		Title:  res.Title,
		ID:     res.ID,
		UserID: res.UserID,
		Alert: dto.Alert{
			Level:   AlertLvlSuccess,
			Message: "create gallery succesfully",
		},
	}, nil
}

func (g *galleryService) ShowGallery(id uint) (dto.ShowGalleryResponse, error) {
	gallery, err := g.Repo.ByID(id)
	if err != nil {
		return dto.ShowGalleryResponse{}, err
	}

	return dto.ShowGalleryResponse{
		Title:  gallery.Title,
		ID:     gallery.ID,
		UserID: gallery.UserID,
	}, nil
}

func 
