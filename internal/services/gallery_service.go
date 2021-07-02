package services

import (
	"tung.gallery/internal/dt/dto"
	"tung.gallery/internal/dt/entity"
	"tung.gallery/internal/repo"
)

type GalleriesServiceInterface interface {
	CreateGallery(*entity.Users, dto.GalleryCreateRequest) (dto.GalleryCreateResponse, error)
	ShowGallery(id uint) (dto.ShowGalleryResponse, error)
	Update(*entity.Users, dto.GalleryUpdateRequest) (dto.GalleryUpdateResponse, error)
	Delete(id uint) (dto.GalleryDeleteResponse, error)
	GetAllGalleriesByUserID(id uint) (dto.UserGetAllGalleriesResponse, error)
}

type galleryService struct {
	Repo repo.GalleryRepositoryInterface
}

func NewGalleryService(r repo.GalleryRepositoryInterface) GalleriesServiceInterface {
	return &galleryService{
		Repo: r,
	}
}

func (g *galleryService) CreateGallery(user *entity.Users, req dto.GalleryCreateRequest) (dto.GalleryCreateResponse, error) {
	gallery := entity.Galleries{
		UserID: user.ID,
		Title:  req.Title,
	}

	res, err := g.Repo.CreateGallery(&gallery)
	if err != nil {
		return dto.GalleryCreateResponse{
			Login: true,
			Title: req.Title,
			Alert: dto.Alert{
				Level:   AlertLvlInfo,
				Message: err.Error(),
			},
		}, err
	}

	return dto.GalleryCreateResponse{
		Login:  true,
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
		return dto.ShowGalleryResponse{Login: true}, err
	}

	return dto.ShowGalleryResponse{
		Login:  true,
		Title:  gallery.Title,
		ID:     gallery.ID,
		UserID: gallery.UserID,
		Alert:  dto.Alert{Level: AlertLvlSuccess, Message: "succes"},
	}, nil
}

func (g *galleryService) Update(user *entity.Users, req dto.GalleryUpdateRequest) (dto.GalleryUpdateResponse, error) {
	gallery, err := g.Repo.ByID(req.ID)
	if err != nil {
		return dto.GalleryUpdateResponse{
			Login: true,
			Title: gallery.Title,
			ID:    gallery.ID,
			Alert: dto.Alert{Level: AlertLvlInfo, Message: err.Error()}}, err
	}
	gallery.Title = req.Title

	err = g.Repo.Update(gallery)
	if err != nil {
		return dto.GalleryUpdateResponse{
			Login: true,
			Title: gallery.Title,
			ID:    gallery.ID,
			Alert: dto.Alert{Level: AlertLvlInfo, Message: err.Error()}}, err
	}

	return dto.GalleryUpdateResponse{
		Title: gallery.Title,
		ID:    gallery.ID,
		Login: true,
		Alert: dto.Alert{Level: AlertLvlSuccess, Message: req.Title}}, nil
}

func (g *galleryService) Delete(id uint) (dto.GalleryDeleteResponse, error) {
	err := g.Repo.Delete(id)
	if err != nil {
		return dto.GalleryDeleteResponse{
			Login: true,
			Alert: dto.Alert{Level: AlertLvlInfo, Message: "success delete gallery"},
		}, err
	}
	return dto.GalleryDeleteResponse{
		Login: true,
		Alert: dto.Alert{Level: AlertLvlSuccess, Message: "success delete gallery"},
	}, nil
}

func (g *galleryService) GetAllGalleriesByUserID(id uint) (dto.UserGetAllGalleriesResponse, error) {
	galleries := make([]dto.Gallrery, 0)

	allGalleries, err := g.Repo.ByUserID(id)
	if err != nil {
		return dto.UserGetAllGalleriesResponse{
			Login: true,
			Alert: dto.Alert{Level: AlertLvlInfo, Message: err.Error()}}, err
	}

	for _, g := range allGalleries {
		gallery := dto.Gallrery{
			Title:  g.Title,
			ID:     g.ID,
			UserID: g.UserID,
		}
		galleries = append(galleries, gallery)
	}

	return dto.UserGetAllGalleriesResponse{
		Galleries: galleries,
		Login:     true,
		Alert: dto.Alert{
			Level:   AlertLvlSuccess,
			Message: "success",
		},
	}, nil
}
