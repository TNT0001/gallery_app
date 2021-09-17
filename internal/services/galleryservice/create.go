package galleryservice

import (
	"errors"
	"log"
	"tung.gallery/internal/dt/dto/gallerydto"
	"tung.gallery/internal/dt/entity"
)

func (g *galleryService) CreateGallery(userID int64, req *gallerydto.GalleryCreateRequest) (
	*gallerydto.GalleryCreateResponse, error) {
	gallery := entity.Galleries{
		UserID: userID,
		Title:  req.Title,
	}

	res, err := g.Repo.CreateGallery(&gallery)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("fail to create gallery")
	}

	return &gallerydto.GalleryCreateResponse{
		Title:  res.Title,
		ID:     res.ID,
		UserID: uint(res.UserID),
	}, nil
}
