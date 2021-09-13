package users

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log"
	"tung.gallery/internal/dt/dto/user_dto"
	"tung.gallery/internal/dt/entity"
	"tung.gallery/pkg/models"
)

func (s *userService) CreateUser(req *user_dto.UserCreateRequest) (*user_dto.UserCreateResponse, error) {
	user, err := s.Repo.ByEmail(req.Email)

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

	res := &user_dto.UserCreateResponse{
		Username:     req.Username,
		Email:        req.Email,
		Password:     req.Password}

	return res, nil
}
