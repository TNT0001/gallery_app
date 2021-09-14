package comment

import (
	"context"
	"errors"
	"log"
	"tung.gallery/internal/dt/dto/comment_dto"
)

func (s *commentService) GetCommentByID(ctx context.Context, commentID int64) (*comment_dto.GetSingleCommentResponse, error){
	comment, err := s.CommentRepo.GetByID(ctx, commentID)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("error when edit comment info")
	}

	return &comment_dto.GetSingleCommentResponse{
		Comment: comment.Comment,
		ImageID: comment.ImageID,
		UserID: comment.UserID,
	}, nil
}

func (s *commentService) GetCommentsByUserID(ctx context.Context, userID int64) (*comment_dto.GetCommentByUserID, error) {
	user, err := s.UserRepo.ByID(userID)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("error when get user info")
	}

	comments, err := s.CommentRepo.ListCommentByUserID(ctx, userID)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("error when get user comments")
	}

	listComment := make([]*comment_dto.GetSingleCommentResponse, 0)
	for _, c := range comments {
		listComment = append(listComment, &comment_dto.GetSingleCommentResponse{
			Comment: c.Comment,
			ImageID: c.ImageID,
			UserID:  c.UserID,
		})
	}
	return &comment_dto.GetCommentByUserID{
		UserID:      userID,
		UserName:    user.Username,
		ListComment: listComment,
	}, nil
}

func (s *commentService) GetCommentsByImageID(ctx context.Context, imageID int64) (*comment_dto.GetCommentByImageID, error) {
	comments, err := s.CommentRepo.ListCommentByimageID(ctx, imageID)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("error when get user comments")
	}

	listComment := make([]*comment_dto.GetSingleCommentResponse, 0)
	for _, c := range comments {
		listComment = append(listComment, &comment_dto.GetSingleCommentResponse{
			Comment: c.Comment,
			ImageID: c.ImageID,
			UserID:  c.UserID,
		})
	}
	return &comment_dto.GetCommentByImageID{
		ImageID: imageID,
		ListComment: listComment,
	}, nil
}
