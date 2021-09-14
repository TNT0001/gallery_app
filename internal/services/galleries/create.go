package galleries

import (
	"errors"
	"log"
	"tung.gallery/internal/dt/dto/gallery_dto"
	"tung.gallery/internal/dt/entity"
)

func (g *galleryService) CreateGallery(userID uint, req *gallery_dto.GalleryCreateRequest) (
	*gallery_dto.GalleryCreateResponse, error) {
	gallery := entity.Galleries{
		UserID: userID,
		Title:  req.Title,
	}

	res, err := g.Repo.CreateGallery(&gallery)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("fail to create gallery")
	}

	return &gallery_dto.GalleryCreateResponse{
		Title:        res.Title,
		ID:           res.ID,
		UserID:       res.UserID,
	}, nil
}
