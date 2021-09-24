package reactrepo

import (
	"context"
	"gorm.io/gorm"
	"tung.gallery/internal/dt/dto/reactdto"
	"tung.gallery/internal/dt/entity"
)

type reactRepo struct {
	DB *gorm.DB
}

func NewReactRepo(db *gorm.DB) ReactRepositoryInterface {
	return &reactRepo{DB: db}
}

func (r *reactRepo) CreateReact(ctx context.Context, react *entity.React) error {
	err := r.DB.WithContext(ctx).Create(react).Error
	return err
}

func (r *reactRepo) ListReactByUserID(ctx context.Context, id int64) ([]*entity.React, error) {
	reacts := make([]*entity.React, 0)
	err := r.DB.WithContext(ctx).Where("user_id = ?", id).Find(&reacts).Error
	if err != nil {
		return nil, err
	}
	return reacts, err
}

func (r *reactRepo) ListReactByImageID(ctx context.Context, id int64) ([]*entity.React, error) {
	reacts := make([]*entity.React, 0)
	err := r.DB.WithContext(ctx).Where("image_id = ?", id).Find(&reacts).Error
	if err != nil {
		return nil, err
	}
	return reacts, err
}

func (r *reactRepo) GetReactCountTotalByImageID(ctx context.Context, id int64) (int64, error) {
	count := new(int64)
	err := r.DB.WithContext(ctx).Model(&entity.React{}).Where("image_id = ?", id).Count(count).Error
	if err != nil {
		return 0, err
	}
	return *count, err
}

func (r *reactRepo) GetReactCountTotalByImageIDEachType(ctx context.Context, id int64) ([]*reactdto.TotalReactIDCountByType, error) {
	totalReact := make([]*reactdto.TotalReactIDCountByType, 0)
	err := r.DB.Model(&entity.React{}).Model(&entity.React{}).WithContext(ctx).Select("type_id, count(id) as total").Where("image_id = ?", id).
		Group("type_id").Find(&totalReact).Error
	if err != nil {
		return nil, err
	}
	return totalReact, err
}

func (r *reactRepo) GetReactByID(ctx context.Context, id int64) (*entity.React, error) {
	react := &entity.React{}
	err := r.DB.WithContext(ctx).Where("id = ?", id).First(react).Error
	if err != nil {
		return nil, err
	}
	return react, err
}

func (r *reactRepo) DeleteReact(ctx context.Context, id int64) error {
	react := &entity.React{}
	react.ID = uint(id)
	err := r.DB.WithContext(ctx).Unscoped().Delete(react).Error
	return err
}

func (r *reactRepo) GetTypeReact(ctx context.Context) ([]*entity.ReactType, error) {
	reactTypes := make([]*entity.ReactType, 0)
	err := r.DB.WithContext(ctx).Find(&reactTypes).Error
	if err != nil {
		return nil, err
	}
	return reactTypes, nil
}

func (r *reactRepo) DeleteReactByUserAndImageID(ctx context.Context, userID, imageID int64) error {
	return r.DB.Unscoped().WithContext(ctx).Where(`image_id = ? and user_id = ?`, imageID, userID).
		Delete(&entity.React{}).Error
}
