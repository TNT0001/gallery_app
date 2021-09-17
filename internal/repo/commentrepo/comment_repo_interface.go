package commentrepo

import (
	"context"
	"tung.gallery/internal/dt/entity"
)

type CommentRepoInterface interface {
	CreateComment(ctx context.Context, comment *entity.Comments) error
	ListCommentByUserID(ctx context.Context, id int64) ([]*entity.Comments, error)
	ListCommentByImageID(ctx context.Context, id int64) ([]*entity.Comments, error)
	GetCommentByID(ctx context.Context, id int64) (*entity.Comments, error)
	EditComment(ctx context.Context, comment *entity.Comments) error
	DeleteComment(ctx context.Context, id int64) error
}
