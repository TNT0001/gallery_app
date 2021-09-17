package commentservice

import (
	"context"
	"errors"
	"log"
	"tung.gallery/internal/dt/dto/commentdto"
)

func (s *commentService) Edit(ctx context.Context, userID int64, req *commentdto.CommentEditRequest) error {
	comment, err := s.Repo.GetCommentByID(ctx, req.CommentID)
	if err != nil {
		log.Println(err.Error())
		return errors.New("error when get comment info")
	}

	if comment.UserID != userID {
		return errors.New("you can't change comment of another")
	}

	comment.Comment = req.Comment
	err = s.Repo.EditComment(ctx, comment)
	if err != nil {
		log.Println(err.Error())
		return errors.New("error when edit comment info")
	}

	return nil
}
