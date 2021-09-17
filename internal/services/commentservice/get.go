package commentservice

import (
	"context"
	"errors"
	"log"
	"tung.gallery/internal/dt/dto/commentdto"
)

func (s *commentService) GetCommentByID(ctx context.Context, commentID int64) (*commentdto.GetSingleCommentResponse, error) {
	comment, err := s.Repo.GetCommentByID(ctx, commentID)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("error when edit comment info")
	}

	return &commentdto.GetSingleCommentResponse{
		Comment: comment.Comment,
		ImageID: comment.ImageID,
		UserID:  comment.UserID,
	}, nil
}

func (s *commentService) GetCommentsByUserID(ctx context.Context, userID int64) (*commentdto.GetCommentByUserID, error) {
	user, err := s.Repo.GetUserByID(userID)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("error when get user info")
	}

	comments, err := s.Repo.ListCommentByUserID(ctx, userID)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("error when get user comments")
	}

	listComment := make([]*commentdto.GetSingleCommentResponse, 0)
	for _, c := range comments {
		listComment = append(listComment, &commentdto.GetSingleCommentResponse{
			Comment: c.Comment,
			ImageID: c.ImageID,
			UserID:  c.UserID,
		})
	}
	return &commentdto.GetCommentByUserID{
		UserID:      userID,
		UserName:    user.Username,
		ListComment: listComment,
	}, nil
}

func (s *commentService) GetCommentsByImageID(ctx context.Context, imageID int64) (*commentdto.GetCommentByImageID, error) {
	comments, err := s.Repo.ListCommentByImageID(ctx, imageID)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("error when get user comments")
	}

	listComment := make([]*commentdto.GetSingleCommentResponse, 0)
	for _, c := range comments {
		listComment = append(listComment, &commentdto.GetSingleCommentResponse{
			Comment: c.Comment,
			ImageID: c.ImageID,
			UserID:  c.UserID,
		})
	}
	return &commentdto.GetCommentByImageID{
		ImageID:     imageID,
		ListComment: listComment,
	}, nil
}
