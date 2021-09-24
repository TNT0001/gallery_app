package commentservice

import (
	"context"
	"tung.gallery/internal/dt/dto/commentdto"
	"tung.gallery/internal/repo"
)

type CommentServiceInterface interface {
	Create(ctx context.Context, request *commentdto.CommentCreateRequest) error
	Delete(ctx context.Context, userID, commentID int64) error
	Edit(ctx context.Context, userID int64, req *commentdto.CommentEditRequest) error
	GetCommentsByImageID(ctx context.Context, imageID int64) (*commentdto.GetCommentByImageID, error)
	GetCommentByID(ctx context.Context, commentID int64) (*commentdto.GetSingleCommentResponse, error)
	GetCommentsByUserID(ctx context.Context, userID int64) (*commentdto.GetCommentByUserID, error)
}

type commentService struct {
	Repo repo.Repo
}

func NewCommentService(r repo.Repo) *commentService {
	return &commentService{
		r,
	}
}
