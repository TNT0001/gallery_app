package comment_repo

import (
	"context"
	"tung.gallery/internal/dt/entity"
)

type CommentRepoInterface interface {
	CreateComment(ctx context.Context, comment *entity.Comments) error
	ListCommentByUserID(ctx context.Context, id int64) ([]*entity.Comments, error)
	ListCommentByimageID(ctx context.Context, id int64) ([]*entity.Comments, error)
	GetByID(ctx context.Context, id int64) (*entity.Comments, error)
	Edit(ctx context.Context, comment *entity.Comments) error
	Delete(ctx context.Context, id int64) error
}
