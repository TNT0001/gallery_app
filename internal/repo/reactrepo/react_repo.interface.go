package reactrepo

import (
	"context"
	"tung.gallery/internal/dt/dto/reactdto"
	"tung.gallery/internal/dt/entity"
)

type ReactRepositoryInterface interface {
	CreateReact(ctx context.Context, react *entity.React) error
	ListReactByUserID(ctx context.Context, id int64) ([]*entity.React, error)
	ListReactByImageID(ctx context.Context, id int64) ([]*entity.React, error)
	GetReactCountTotalByImageID(ctx context.Context, id int64) (int64, error)
	GetReactCountTotalByImageIDEachType(ctx context.Context, id int64) ([]*reactdto.TotalReactIDCountByType, error)
	GetReactByID(ctx context.Context, id int64) (*entity.React, error)
	DeleteReact(ctx context.Context, id int64) error
	GetTypeReact(ctx context.Context) ([]*entity.ReactType, error)
}
