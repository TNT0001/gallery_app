package reactsservice

import (
	"context"
	"errors"
	"log"
)

func (s *reactService) DeleteReact(ctx context.Context, userID, reactID int64) error {
	reacts, err := s.Repo.GetReactByID(ctx, reactID)
	if err != nil {
		log.Println(err.Error())
		return errors.New("error when get react info")
	}

	if reacts.UserID != userID {
		return errors.New("you don't have permission to delete react create by others")
	}

	err = s.Repo.DeleteReact(ctx, reactID)
	if err != nil {
		log.Println(err.Error())
		return errors.New("error when delete react")
	}

	return nil
}
