package reactsservice

import (
	"context"
	"errors"
	"log"
	"tung.gallery/internal/dt/dto/reactdto"
	"tung.gallery/internal/dt/entity"
	"tung.gallery/internal/repo"
)

func (s *reactService) Create(ctx context.Context, req *reactdto.CreateReactRequest) error {
	err := s.Repo.Transactions(func(subRepo repo.Repo) error {
		err := subRepo.DeleteReactByUserAndImageID(ctx, req.UserID, req.ImageID)
		if err != nil {
			return err
		}

		err = subRepo.CreateReact(ctx, &entity.React{
			TypeID:  req.TypeID,
			UserID:  req.UserID,
			ImageID: req.ImageID,
		})
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		log.Println(err.Error())
		return errors.New("error when create react")
	}

	return nil
}
