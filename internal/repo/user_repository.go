package repo

import (
	"gorm.io/gorm"
	"tung.gallery/internal/dt/entity"
	"tung.gallery/pkg/models"
)

type UserRepositoryInterface interface {
	CreateUser(user entity.Users) error
	ByEmail(email string) (*entity.Users, error)
	ByID(id uint) (*entity.Users, error)
	Update(user entity.Users) error
	Delete(id uint) error
}

type userRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepositoryInterface {
	return &userRepo{DB: db}
}

func (u *userRepo) CreateUser(user entity.Users) error {
	err := u.DB.Create(&user).Error
	return err
}

func (u *userRepo) ByEmail(email string) (*entity.Users, error) {
	user := entity.Users{}
	err := u.DB.First(&user, "email=?", email).Error
	if err == gorm.ErrRecordNotFound {
		return &user, models.ErrNotFound
	}
	return &user, err
}

func (u *userRepo) ByID(id uint) (*entity.Users, error) {
	user := entity.Users{}
	err := u.DB.First(&user, id).Error
	if err == gorm.ErrRecordNotFound {
		return nil, models.ErrNotFound
	}
	return &user, err
}

func (u *userRepo) Update(user entity.Users) error {
	err := u.DB.Updates(&user).Error
	return err
}

func (u *userRepo) Delete(id uint) error {
	if id < 0 {
		return models.ErrInvalidID
	}
	user := entity.Users{}
	user.ID = id
	err := u.DB.First(&user).Error
	if err != nil {
		return err
	}

	err = u.DB.Delete(user).Error
	return err
}
