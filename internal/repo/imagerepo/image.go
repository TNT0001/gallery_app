package imagerepo

import (
	"context"
	"tung.gallery/internal/dt/dto/imagedto"
	"tung.gallery/internal/dt/entity"
)

func (r *imageRepository) GetImageByID(ctx context.Context, id int64) (*entity.Images, error) {
	image := &entity.Images{}
	err := r.db.WithContext(ctx).Where("id = ?", id).First(image).Error
	if err != nil {
		return nil, err
	}
	return image, nil
}

func (r *imageRepository) GetImagesByGalleryID(ctx context.Context, galleryID []int64) ([]*entity.Images, error) {
	image := make([]*entity.Images, 0)
	err := r.db.WithContext(ctx).Where("gallery_id in ?", galleryID).Find(&image).Error
	if err != nil {
		return nil, err
	}
	return image, nil
}

func (r *imageRepository) CreateImage(ctx context.Context, req *imagedto.ImageUploadRequest) error {
	image := &entity.Images{
		UserID:    req.UserID,
		GalleryID: req.GalleryID,
		ImageURL:  req.ImageURL,
		Title:     req.Title,
		ImageUUID: req.ImageUUID,
	}

	return r.db.WithContext(ctx).Create(image).Error
}
