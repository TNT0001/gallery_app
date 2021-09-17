package userservice

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log"
	"tung.gallery/internal/dt/dto/userdto"
	"tung.gallery/internal/dt/entity"
	"tung.gallery/pkg/models"
)

func (s *userService) FindUserByID(id int64) (*entity.Users, error) {
	user, err := s.Repo.GetUserByID(id)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("fail to find user by id")
	}
	return user, nil
}

func (s *userService) FindUserByEmail(req *userdto.UserLoginRequest) (*entity.Users, error) {
	user, err := s.Repo.GetUserByEmail(req.Email)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("fail to find user by email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			log.Println(err.Error())
			return nil, models.ErrInvalidPassword
		} else {
			log.Println(err.Error())
			return nil, errors.New("fail to  check password")
		}
	}

	return user, nil
}
