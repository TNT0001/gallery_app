package reactsservice

import (
	"context"
	"tung.gallery/internal/dt/dto/reactdto"
	"tung.gallery/internal/repo"
)

type ReactServiceInterface interface {
	Create(ctx context.Context, req *reactdto.CreateReactRequest) error
	DeleteReact(ctx context.Context, userID, reactID int64) error
	GetReactByID(ctx context.Context, reactID int64) (*reactdto.GetReactByIDResponse, error)
	GetReactByUserID(ctx context.Context, userID int64) (*reactdto.GetListReactByUserIDResponse, error)
	GetReactByImageID(ctx context.Context, imageID int64) (*reactdto.GetListReactByImagesIDResponse, error)
	GetReactCountByImageID(ctx context.Context, imageID int64) (*reactdto.GetReactTotalByImageIDResponse, error)
	GetReactCountByEachTypeAndImageID(ctx context.Context, imageID int64) (*reactdto.GetReactTotalByEachTypeAndImageIDResponse, error)
	GetReactMap(ctx context.Context) (map[int64]string, error)
}

type reactService struct {
	Repo repo.GalleryRepository
}

func NewReactService(r repo.GalleryRepository) ReactServiceInterface {
	return &reactService{
		Repo: r,
	}
}
