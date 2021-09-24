package galleryrepo

import (
	"tung.gallery/internal/dt/entity"
)

func (u *galleryRepo) CreateGallery(gallery *entity.Galleries) (*entity.Galleries, error) {
	err := u.DB.Create(gallery).Error
	return gallery, err
}

func (u *galleryRepo) GetGalleryByID(id int64) (*entity.Galleries, error) {
	gallery := &entity.Galleries{}
	gallery.ID = uint(id)

	err := u.DB.First(gallery).Error
	if err != nil {
		return nil, err
	}

	return gallery, nil
}

func (u *galleryRepo) GetGalleriesByListID(id []int64) ([]*entity.Galleries, error) {
	gallery := make([]*entity.Galleries, 0)

	err := u.DB.Where("id in ?", id).Find(&gallery).Error
	if err != nil {
		return nil, err
	}

	return gallery, nil
}

func (u *galleryRepo) UpdateGallery(gallery *entity.Galleries) error {
	err := u.DB.Updates(gallery).Error
	return err
}

func (u *galleryRepo) DeleteGallery(id int64) error {
	gallery := &entity.Galleries{}
	gallery.ID = uint(id)
	err := u.DB.Unscoped().Delete(gallery).Error
	return err
}

func (u *galleryRepo) GetGalleryByUserID(userID int64) ([]*entity.Galleries, error) {
	galleries := make([]*entity.Galleries, 0)

	err := u.DB.Find(&galleries, "galleries.user_id =?", userID).Error
	if err != nil {
		return nil, err
	}

	return galleries, nil
}
