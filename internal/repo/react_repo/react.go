package react_repo

import (
	"context"
	"gorm.io/gorm"
	"tung.gallery/internal/dt/dto/react_dto"
	"tung.gallery/internal/dt/entity"
)

type reactRepo struct {
	DB *gorm.DB
}

func NewReactRepo(db *gorm.DB) *reactRepo {
	return &reactRepo{DB: db}
}

func (r *reactRepo) CreateReact(ctx context.Context, react *entity.React) error {
	err := r.DB.WithContext(ctx).Create(react).Error
	return err
}

func (r *reactRepo) ListReactByUserID(ctx context.Context, id uint) ([]*entity.React, error) {
	reacts := make([]*entity.React, 0)
	err := r.DB.WithContext(ctx).Where("user_id = ?", id).Find(&reacts).Error
	if err != nil {
		return nil, err
	}
	return reacts, err
}

func (r *reactRepo) ListReactByImageID(ctx context.Context, id uint) ([]*entity.React, error) {
	reacts := make([]*entity.React, 0)
	err := r.DB.WithContext(ctx).Where("image_id = ?", id).Find(&reacts).Error
	if err != nil {
		return nil, err
	}
	return reacts, err
}

func (r *reactRepo) GetCountTotalByImageID(ctx context.Context, id uint) (*int64, error) {
	count := new(int64)
	err := r.DB.WithContext(ctx).Where("image_id = ?", id).Count(count).Error
	if err != nil {
		return nil, err
	}
	return count, err
}

func (r *reactRepo) GetCountTotalByImageIDEachType(ctx context.Context, id uint) (*react_dto.TotalReactCountByType, error) {
	totalReact := &react_dto.TotalReactCountByType{}
	err := r.DB.Model(&entity.React{}).WithContext(ctx).Select("type_id, count(id) as total").Where("image_id = ?", id).
		Group("type_id").Find(totalReact).Error
	if err != nil {
		return nil, err
	}
	return totalReact, err
}

func (r *reactRepo) GetByID(ctx context.Context, id uint) (*entity.React, error) {
	react := &entity.React{}
	err := r.DB.WithContext(ctx).Where("id = ?", id).First(react).Error
	if err != nil {
		return nil, err
	}
	return react, err
}

func (r *reactRepo) Delete(ctx context.Context, id uint) error {
	react := &entity.React{}
	react.ID = id
	err := r.DB.WithContext(ctx).Unscoped().Delete(react).Error
	return err
}

func (r *reactRepo) GetTypeReact (ctx context.Context) ([]*entity.ReactType, error) {
	reactTypes := make([]*entity.ReactType, 0)
	err := r.DB.WithContext(ctx).Find(&reactTypes).Error
	if err != nil {
		return nil, err
	}
	return reactTypes, nil
}
