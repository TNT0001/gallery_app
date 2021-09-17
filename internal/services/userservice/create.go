package userservice

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log"
	"tung.gallery/internal/dt/dto/userdto"
	"tung.gallery/internal/dt/entity"
	"tung.gallery/pkg/models"
)

func (s *userService) CreateUser(req *userdto.UserCreateRequest) (*userdto.UserCreateResponse, error) {
	user, err := s.Repo.GetUserByEmail(req.Email)

	if err != nil && err != models.ErrNotFound {
		log.Println(err.Error())
		return nil, models.ErrInternalServerError
	}

	if user.Email != "" {
		return nil, models.ErrEmailHasExist
	}

	newUser := &entity.Users{
		Username: req.Username,
		Email:    req.Email,
	}

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	newUser.Password = string(hashPassword)
	err = s.Repo.CreateUser(newUser)

	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("error when create user")
	}

	res := &userdto.UserCreateResponse{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password}

	return res, nil
}
