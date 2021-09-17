package commentservice

import (
	"context"
	"errors"
	"log"
)

func (s *commentService) Delete(ctx context.Context, userID, commentID int64) error {
	comment, err := s.Repo.GetCommentByID(ctx, commentID)
	if err != nil {
		log.Println(err.Error())
		return errors.New("error when get comment info")
	}

	if comment.UserID != userID {
		return errors.New("you can't delete comment of another")
	}

	err = s.Repo.DeleteComment(ctx, commentID)
	if err != nil {
		log.Println(err.Error())
		return errors.New("error when delete comment")
	}

	return nil
}
