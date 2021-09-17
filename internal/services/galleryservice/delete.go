package galleryservice

import (
	"errors"
	"log"
	"tung.gallery/internal/dt/dto/gallerydto"
)

func (g *galleryService) Delete(userID, galleryID int64) (*gallerydto.GalleryDeleteResponse, error) {
	gallery, err := g.Repo.GetGalleryByID(galleryID)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("error when get gallery info")
	}

	if gallery.UserID != userID {
		return nil, errors.New("you don't have permission to delete this gallery")
	}

	err = g.Repo.DeleteGallery(galleryID)
	if err != nil {
		return nil, errors.New("fail to delete gallery, please try again later")
	}

	return &gallerydto.GalleryDeleteResponse{}, nil
}
