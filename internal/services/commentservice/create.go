package commentservice

import (
	"context"
	"errors"
	"log"
	"tung.gallery/internal/dt/dto/commentdto"
	"tung.gallery/internal/dt/entity"
)

func (s *commentService) Create(ctx context.Context, req *commentdto.CommentCreateRequest) error {
	image, err := s.Repo.GetImageByID(ctx, req.ImageID)
	if err != nil {
		log.Println(err.Error())
		return errors.New("error when get image info")
	}

	gallery, err := s.Repo.GetGalleryByID(image.GalleryID)
	if err != nil {
		log.Println(err.Error())
		return errors.New("error when get gallery info")
	}

	if !gallery.IsPublic && gallery.UserID != req.UserID {
		return errors.New("you don't have permission to comment on this image")
	}

	commentEntity := &entity.Comments{
		Comment: req.Comment,
		UserID:  req.UserID,
		ImageID: req.ImageID,
	}

	err = s.Repo.CreateComment(ctx, commentEntity)
	if err != nil {
		log.Println(err.Error())
		return errors.New("error when create comment")
	}
	return nil
}
