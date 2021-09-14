package comment

import (
	"context"
	"errors"
	"log"
	"tung.gallery/internal/dt/dto/comment_dto"
)

func (s *commentService) Edit(ctx context.Context, userID int64, req *comment_dto.CommentEditRequest) error{
	comment, err := s.CommentRepo.GetByID(ctx, req.CommentID)
	if err != nil {
		log.Println(err.Error())
		return errors.New("error when get comment info")
	}

	if comment.UserID != userID {
		return errors.New("you can't change comment of another")
	}

	comment.Comment = req.Comment
	err = s.CommentRepo.Edit(ctx, comment)
	if err != nil {
		log.Println(err.Error())
		return errors.New("error when edit comment info")
	}

	return nil
}
