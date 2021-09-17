package imageservice

import (
	"context"
	"errors"
	"fmt"
	uuid "github.com/nu7hatch/gouuid"
	"log"
	"os"
	"tung.gallery/internal/dt/dto/imagedto"
)

func (s *imageService) CreateImage(ctx context.Context, req *imagedto.ImageUploadRequest, content string) (string, error) {
	imageUUID, err := uuid.NewV4()
	if err != nil {
		log.Println(err.Error())
		return "", errors.New("error when get create uuid for image")
	}

	gallery, err := s.Repo.GetGalleryByID(req.GalleryID)
	if err != nil {
		log.Println(err.Error())
		return "", errors.New("error when get info of gallery")
	}

	if req.UserID != gallery.UserID {
		return "", errors.New("you dont have permission")
	}

	imageURL, err := s.StoredImage(req.GalleryID, req.UserID, imageUUID.String(), content)
	if err != nil {
		log.Println(err.Error())
		return "", errors.New("error when stored image")
	}

	req.ImageURL = imageURL
	req.ImageUUID = imageUUID.String()

	err = s.Repo.CreateImage(ctx, req)
	if err != nil {
		log.Println(err.Error())
		return "", errors.New("error when create image")
	}

	return imageURL, nil
}

func (s *imageService) StoredImage(galleryID, userID int64, imageUID, content string) (string, error) {
	pathPrefix := "../../assets/images/"
	dir := pathPrefix + fmt.Sprintf("user_%d/gallery_%d", userID, galleryID)
	url := fmt.Sprintf("user_%d/gallery_%d/%s.txt", userID, galleryID, imageUID)
	filePath := pathPrefix + url
	err := os.MkdirAll(dir, 0777)
	if err != nil {
		return "", err
	}

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		return "", err
	}

	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return "", err
	}

	return url, nil
}
