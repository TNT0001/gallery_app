package user

import (
	"gorm.io/gorm"
	"tung.gallery/internal/dt/entity"
	"tung.gallery/internal/repo/userrepo"
	"tung.gallery/pkg/models"
)

type userRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) userrepo.UserRepositoryInterface {
	return &userRepo{DB: db}
}

func (r *userRepo) CreateUser(user *entity.Users) error {
	err := r.DB.Create(user).Error
	return err
}

func (r *userRepo) GetUserByEmail(email string) (*entity.Users, error) {
	user := entity.Users{}
	err := r.DB.First(&user, "email=?", email).Error
	if err == gorm.ErrRecordNotFound {
		return &user, models.ErrNotFound
	}
	return &user, err
}

func (r *userRepo) GetUserByID(id int64) (*entity.Users, error) {
	user := entity.Users{}
	err := r.DB.First(&user, id).Error
	if err == gorm.ErrRecordNotFound {
		return nil, models.ErrNotFound
	}
	return &user, err
}

func (r *userRepo) UpdateUser(user *entity.Users, id int64) error {
	err := r.DB.Where("id = ?", id).Updates(&user).Error
	return err
}

func (r *userRepo) DeleteUser(id int64) error {
	if id < 0 {
		return models.ErrInvalidID
	}
	user := &entity.Users{}
	user.ID = uint(id)
	err := r.DB.First(&user).Error
	if err != nil {
		return err
	}

	err = r.DB.Unscoped().Delete(user).Error
	return err
}

func (r *userRepo) GetListUserByListID(id []int64) ([]*entity.Users, error) {
	user := make([]*entity.Users, 0)
	err := r.DB.Where("id in ?", id).Find(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, models.ErrNotFound
	}
	return user, err
}
