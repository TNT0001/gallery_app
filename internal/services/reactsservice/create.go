package reactsservice

import (
	"context"
	"errors"
	"log"
	"tung.gallery/internal/dt/dto/reactdto"
	"tung.gallery/internal/dt/entity"
)

func (s *reactService) Create(ctx context.Context, req *reactdto.CreateReactRequest) error {
	err := s.Repo.CreateReact(ctx, &entity.React{
		TypeID:  req.TypeID,
		UserID:  req.UserID,
		ImageID: req.ImageID,
	})

	if err != nil {
		log.Println(err.Error())
		return errors.New("error when create react")
	}

	return nil
}
