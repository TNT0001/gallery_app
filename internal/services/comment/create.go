package comment

import (
	"context"
	"errors"
	"log"
	"tung.gallery/internal/dt/entity"
)

func (s *commentService) Create(ctx context.Context, userID, imageID int64, comment string) error{
	image, err := s.ImageRepo.GetImageByID(ctx, uint(imageID))
	if err != nil {
		log.Println(err.Error())
		return errors.New("error when get image info")
	}

	gallery, err := s.GalleryRepo.ByID(image.GalleryID)
	if err != nil {
		log.Println(err.Error())
		return errors.New("error when get gallery info")
	}

	if !gallery.IsPublic && gallery.UserID != uint(userID) {
		return errors.New("you don't have permission to comment on this image")
	}

	commentEntity := &entity.Comments{
		Comment: comment,
		UserID:  userID,
		ImageID: imageID,
	}

	err = s.CommentRepo.CreateComment(ctx, commentEntity)
	if err != nil {
		log.Println(err.Error())
		return errors.New("error when create comment")
	}
	return nil
}
