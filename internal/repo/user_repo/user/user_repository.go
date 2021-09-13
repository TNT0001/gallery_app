package user

import (
	"gorm.io/gorm"
	"tung.gallery/internal/dt/entity"
	"tung.gallery/internal/repo/user_repo"
	"tung.gallery/pkg/models"
)

type userRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) user_repo.UserRepositoryInterface {
	return &userRepo{DB: db}
}

func (r *userRepo) CreateUser(user *entity.Users) error {
	err := r.DB.Create(user).Error
	return err
}

func (r *userRepo) ByEmail(email string) (*entity.Users, error) {
	user := entity.Users{}
	err := r.DB.First(&user, "email=?", email).Error
	if err == gorm.ErrRecordNotFound {
		return &user, models.ErrNotFound
	}
	return &user, err
}

func (r *userRepo) ByID(id uint) (*entity.Users, error) {
	user := entity.Users{}
	err := r.DB.First(&user, id).Error
	if err == gorm.ErrRecordNotFound {
		return nil, models.ErrNotFound
	}
	return &user, err
}

func (r *userRepo) Update(user *entity.Users, id uint) error {
	err := r.DB.Where("id = ?", id).Updates(&user).Error
	return err
}

func (r *userRepo) Delete(id uint) error {
	if id < 0 {
		return models.ErrInvalidID
	}
	user := &entity.Users{}
	user.ID = id
	err := r.DB.First(&user).Error
	if err != nil {
		return err
	}

	err = r.DB.Unscoped().Delete(user).Error
	return err
}

func (r *userRepo) ByListID(id []uint) ([]*entity.Users, error) {
	user := make([]*entity.Users, 0)
	err := r.DB.Where("id in ?", id).Find(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, models.ErrNotFound
	}
	return user, err
}
