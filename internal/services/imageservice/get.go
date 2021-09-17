package imageservice

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"tung.gallery/internal/dt/dto/imagedto"
)

func (s *imageService) GetImageByID(ctx context.Context, userID, imageID int64) (*imagedto.GetImageResponse, error) {
	image, err := s.Repo.GetImageByID(ctx, imageID)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("error when get image")
	}

	gallery, err := s.Repo.GetGalleryByID(image.GalleryID)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("error when get gallery info")
	}

	if !gallery.IsPublic && gallery.UserID != userID {
		return nil, errors.New("you don't have permission to get this image")
	}

	content, err := s.GetContentImage(image.ImageURL)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("error when get content image")
	}

	return &imagedto.GetImageResponse{
		UserID:    image.UserID,
		GalleryID: image.GalleryID,
		ImageURL:  image.ImageURL,
		Title:     image.Title,
		Content:   content,
	}, nil
}

func (s *imageService) GetImageByGalleryID(ctx context.Context, userID int64, galleryID []int64) (*imagedto.GetImageListResponse, error) {
	galleries, err := s.Repo.GetGalleriesByListID(galleryID)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("error when get gallery info")
	}

	allowGallery := make([]int64, 0)
	for _, g := range galleries {
		if g.IsPublic || userID == g.UserID {
			allowGallery = append(allowGallery, int64(g.ID))
		}
	}

	if len(allowGallery) == 0 {
		return nil, errors.New("you don't have permission to get image")
	}

	images, err := s.Repo.GetImagesByGalleryID(ctx, allowGallery)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("error when get image")
	}

	res := make([]*imagedto.GetImageResponse, 0)

	for _, i := range images {
		content, err := s.GetContentImage(i.ImageURL)
		if err != nil {
			log.Println(err.Error())
			return nil, errors.New("error when get content image")
		}

		res = append(res, &imagedto.GetImageResponse{
			UserID:    i.UserID,
			GalleryID: i.GalleryID,
			ImageURL:  i.ImageURL,
			Title:     i.Title,
			Content:   content,
		})
	}

	return &imagedto.GetImageListResponse{
		Images: res,
	}, nil
}

func (s *imageService) GetImageByUserID(ctx context.Context, currentUser, userID int64) (*imagedto.GetImageListResponse, error) {
	galleries, err := s.Repo.GetGalleryByUserID(userID)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("error when get gallery from user")
	}

	galleryList := make([]int64, 0)
	for _, g := range galleries {
		if g.IsPublic || g.UserID == currentUser {
			galleryList = append(galleryList, int64(g.ID))
		}
	}

	if len(galleryList) == 0 {
		return nil, errors.New("you don't have permission to see image of that user")
	}

	images, err := s.Repo.GetImagesByGalleryID(ctx, galleryList)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("error when get image")
	}

	res := make([]*imagedto.GetImageResponse, 0)

	for _, i := range images {
		content, err := s.GetContentImage(i.ImageURL)
		if err != nil {
			log.Println(err.Error())
			return nil, errors.New("error when get content image")
		}

		res = append(res, &imagedto.GetImageResponse{
			UserID:    i.UserID,
			GalleryID: i.GalleryID,
			ImageURL:  i.ImageURL,
			Title:     i.Title,
			Content:   content,
		})
	}

	return &imagedto.GetImageListResponse{
		Images: res,
	}, nil
}

func (s *imageService) GetContentImage(imageURL string) (string, error) {
	filePath := fmt.Sprintf("../../assets/images/%s", imageURL)
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
