package galleries

import (
	"errors"
	"log"
	"tung.gallery/internal/dt/dto/gallery_dto"
)

func (g *galleryService) Delete(userID, galleryID uint) (*gallery_dto.GalleryDeleteResponse, error) {
	gallery, err := g.Repo.ByID(galleryID)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("error when get gallery info")
	}

	if gallery.UserID != userID {
		return nil, errors.New("you don't have permission to delete this gallery")
	}

	err = g.Repo.Delete(galleryID)
	if err != nil {
		return nil, errors.New("fail to delete gallery, please try again later")
	}

	return &gallery_dto.GalleryDeleteResponse{}, nil
}
