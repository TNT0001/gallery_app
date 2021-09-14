package galleries

import (
	"errors"
	"log"
	"tung.gallery/internal/dt/dto/gallery_dto"
)

func (g *galleryService) Update(userID uint, req *gallery_dto.GalleryUpdateRequest) (*gallery_dto.GalleryUpdateResponse, error) {
	gallery, err := g.Repo.ByID(req.ID)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("fail to get gallery")
	}

	if userID != gallery.UserID {
		return nil, errors.New("you don't have permission to update gallery don't belong to you")
	}

	gallery.Title = req.Title

	err = g.Repo.Update(gallery)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("fail to update gallery")
	}

	return &gallery_dto.GalleryUpdateResponse{
		ID:    req.ID,
		Title: req.Title,
	}, nil
}
