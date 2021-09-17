package galleryservice

import (
	"errors"
	"log"
	"tung.gallery/internal/dt/dto/gallerydto"
)

func (g *galleryService) GetAllGalleriesByUserID(currentUserID, userID int64) (*gallerydto.ShowAllGalleryByUserIDResponse, error) {
	res := make([]*gallerydto.ShowGalleryResponse, 0)

	Galleries, err := g.Repo.GetGalleryByUserID(userID)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("fail to get gallery")
	}

	for _, g := range Galleries {
		gallery := &gallerydto.ShowGalleryResponse{
			Title:  g.Title,
			ID:     int64(g.ID),
			UserID: g.UserID,
		}
		if currentUserID == userID || g.IsPublic {
			res = append(res, gallery)
		}
	}

	return &gallerydto.ShowAllGalleryByUserIDResponse{Galleries: res}, nil
}

func (g *galleryService) GetGalleryByID(userID, galleryID int64) (*gallerydto.ShowGalleryResponse, error) {
	gallery, err := g.Repo.GetGalleryByID(galleryID)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("fail to get gallery")
	}

	res := &gallerydto.ShowGalleryResponse{}
	if gallery.UserID == userID || gallery.IsPublic {
		res.Title = gallery.Title
		res.ID = int64(gallery.ID)
		res.UserID = gallery.UserID
		return res, nil
	} else {
		return nil, errors.New("you don't have permission to see this gallery")
	}
}
