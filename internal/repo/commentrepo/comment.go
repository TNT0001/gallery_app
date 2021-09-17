package commentrepo

import (
	"context"
	"gorm.io/gorm"
	"tung.gallery/internal/dt/entity"
	"tung.gallery/pkg/models"
)

type commentRepo struct {
	DB *gorm.DB
}

func NewCommentRepo(db *gorm.DB) CommentRepoInterface {
	return &commentRepo{DB: db}
}

func (r *commentRepo) CreateComment(ctx context.Context, comment *entity.Comments) error {
	err := r.DB.WithContext(ctx).Create(comment).Error
	return err
}

func (r *commentRepo) ListCommentByUserID(ctx context.Context, userID int64) ([]*entity.Comments, error) {
	comments := make([]*entity.Comments, 0)
	err := r.DB.WithContext(ctx).Where("user_id = ?", userID).Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, err
}

func (r *commentRepo) ListCommentByImageID(ctx context.Context, imageID int64) ([]*entity.Comments, error) {
	comments := make([]*entity.Comments, 0)
	err := r.DB.WithContext(ctx).Where("image_id = ?", imageID).Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, err
}

func (r *commentRepo) GetCommentByID(ctx context.Context, commentID int64) (*entity.Comments, error) {
	comment := &entity.Comments{}
	err := r.DB.WithContext(ctx).Where("id = ?", commentID).First(comment).Error
	if err != nil {
		return nil, err
	}
	return comment, err
}

func (r *commentRepo) EditComment(ctx context.Context, comment *entity.Comments) error {
	err := r.DB.WithContext(ctx).Updates(comment).Error
	return err
}

func (r *commentRepo) DeleteComment(ctx context.Context, commentID int64) error {
	if commentID < 0 {
		return models.ErrInvalidID
	}
	comment := &entity.Comments{}
	comment.ID = uint(commentID)
	err := r.DB.WithContext(ctx).Unscoped().Delete(comment).Error
	return err
}
