package galleries

import (
	"errors"
	"log"
	"tung.gallery/internal/dt/dto/gallery_dto"
)

func (g *galleryService) GetAllGalleriesByUserID(currentUserID, userID uint) (*gallery_dto.ShowAllGalleryByUserIDResponse, error) {
	res := make([]*gallery_dto.ShowGalleryResponse, 0)

	Galleries, err := g.Repo.ByUserID(userID)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("fail to get gallery")
	}

	for _, g := range Galleries {
		gallery := &gallery_dto.ShowGalleryResponse{
			Title:  g.Title,
			ID:     g.ID,
			UserID: g.UserID,
		}
		if currentUserID == userID || g.IsPublic {
			res = append(res, gallery)
		}
	}

	return &gallery_dto.ShowAllGalleryByUserIDResponse{Galleries:res}, nil
}

func (g *galleryService) GetGalleryByID(userID, galleryID uint) (*gallery_dto.ShowGalleryResponse, error){
	gallery, err := g.Repo.ByID(galleryID)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("fail to get gallery")
	}

	res := &gallery_dto.ShowGalleryResponse{}
	if gallery.UserID == userID || gallery.IsPublic {
		res.Title = gallery.Title
		res.ID = gallery.ID
		res.UserID = gallery.UserID
		return res, nil
	} else {
		return nil, errors.New("you don't have permission to see this gallery")
	}
}
