package userservice

import (
	"errors"
	"log"
	"tung.gallery/internal/dt/dto/userdto"
	"tung.gallery/internal/dt/entity"
	"tung.gallery/pkg/models"
)

func (s *userService) UpdateUser(oldUser *entity.Users, req *userdto.UserUpdateRequest) (*userdto.UserUpdateResponse, error) {
	if req.Email != "" {
		_, err := s.Repo.GetUserByEmail(req.Email)
		if err != nil && err != models.ErrNotFound {
			return nil, errors.New("error when check user email")
		} else if err == nil {
			return nil, errors.New("email has exist")
		}
	}

	user := &entity.Users{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
		Age:      req.Age,
		Birthday: req.Birthday,
		ImageURL: req.ImageURL,
	}

	err := s.Repo.UpdateUser(user, int64(oldUser.ID))
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("fail to update user profile")
	}

	return &userdto.UserUpdateResponse{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
		Age:      user.Age,
		ImageURL: user.ImageURL,
		Birthday: user.Birthday,
	}, nil
}
