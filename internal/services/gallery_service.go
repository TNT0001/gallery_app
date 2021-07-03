package services

import (
	"tung.gallery/internal/dt/dto"
	"tung.gallery/internal/dt/entity"
	"tung.gallery/internal/repo"
	"tung.gallery/pkg/models"
	"tung.gallery/pkg/utils"
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
		baseResponse := utils.BaseResponse(false, AlertLvlInfo, models.ErrCreateGalleryFail.Error())
		return dto.GalleryCreateResponse{
			BaseResponse: baseResponse,
		}, err
	}

	baseResponse := utils.BaseResponse(false, AlertLvlSuccess, "successful create gallery")
	return dto.GalleryCreateResponse{
		Login:        true,
		Title:        res.Title,
		ID:           res.ID,
		UserID:       res.UserID,
		BaseResponse: baseResponse,
	}, nil
}

func (g *galleryService) ShowGallery(id uint) (dto.ShowGalleryResponse, error) {
	gallery, err := g.Repo.ByID(id)
	if err != nil {
		baseResponse := utils.BaseResponse(false, AlertLvlInfo, models.ErrCreateGalleryFail.Error())
		return dto.ShowGalleryResponse{BaseResponse: baseResponse}, err
	}

	return dto.ShowGalleryResponse{
		Title:        gallery.Title,
		ID:           gallery.ID,
		UserID:       gallery.UserID,
		BaseResponse: dto.BaseResponse{},
	}, nil
}

func (g *galleryService) Update(user *entity.Users, req dto.GalleryUpdateRequest) (dto.GalleryUpdateResponse, error) {
	gallery, err := g.Repo.ByID(req.ID)
	if err != nil {
		baseResponse := utils.BaseResponse(false, AlertLvlInfo, models.ErrShowGalleryFail.Error())
		return dto.GalleryUpdateResponse{
			Title:        gallery.Title,
			ID:           gallery.ID,
			BaseResponse: baseResponse}, err
	}

	gallery.Title = req.Title

	err = g.Repo.Update(gallery)
	if err != nil {
		baseResponse := utils.BaseResponse(false, AlertLvlInfo, models.ErrorUpdateGalleryFail.Error())
		return dto.GalleryUpdateResponse{
			Title:        gallery.Title,
			ID:           gallery.ID,
			BaseResponse: baseResponse}, err
	}

	baseResponse := utils.BaseResponse(false, AlertLvlSuccess, "update successfully")
	return dto.GalleryUpdateResponse{
		Title:        gallery.Title,
		ID:           gallery.ID,
		BaseResponse: baseResponse}, nil
}

func (g *galleryService) Delete(id uint) (dto.GalleryDeleteResponse, error) {
	err := g.Repo.Delete(id)
	if err != nil {
		baseResponse := utils.BaseResponse(false, AlertLvlInfo, models.ErrorDeleteGalleryFail.Error())
		return dto.GalleryDeleteResponse{BaseResponse: baseResponse}, err
	}

	baseResponse := utils.BaseResponse(false, AlertLvlSuccess, "success delete gallery")
	return dto.GalleryDeleteResponse{BaseResponse: baseResponse}, nil
}

func (g *galleryService) GetAllGalleriesByUserID(id uint) (dto.UserGetAllGalleriesResponse, error) {
	galleries := make([]dto.Gallrery, 0)

	allGalleries, err := g.Repo.ByUserID(id)
	if err != nil {
		baseResponse := utils.BaseResponse(false, AlertLvlInfo, models.ErrShowGalleryFail.Error())
		return dto.UserGetAllGalleriesResponse{BaseResponse: baseResponse}, err
	}

	for _, g := range allGalleries {
		gallery := dto.Gallrery{
			Title:  g.Title,
			ID:     g.ID,
			UserID: g.UserID,
		}
		galleries = append(galleries, gallery)
	}

	baseResponse := utils.BaseResponse(false, AlertLvlSuccess, "success get all gallery")
	return dto.UserGetAllGalleriesResponse{
		Galleries:    galleries,
		BaseResponse: baseResponse,
	}, nil
}
