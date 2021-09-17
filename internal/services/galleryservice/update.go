package galleryservice

import (
	"errors"
	"log"
	"tung.gallery/internal/dt/dto/gallerydto"
)

func (g *galleryService) Update(userID int64, req *gallerydto.GalleryUpdateRequest) (*gallerydto.GalleryUpdateResponse, error) {
	gallery, err := g.Repo.GetGalleryByID(req.ID)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("fail to get gallery")
	}

	if userID != gallery.UserID {
		return nil, errors.New("you don't have permission to update gallery don't belong to you")
	}

	gallery.Title = req.Title

	err = g.Repo.UpdateGallery(gallery)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("fail to update gallery")
	}

	return &gallerydto.GalleryUpdateResponse{
		ID:    req.ID,
		Title: req.Title,
	}, nil
}
