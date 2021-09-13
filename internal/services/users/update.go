package users

import (
	"errors"
	"log"
	"tung.gallery/internal/dt/dto/user_dto"
	"tung.gallery/internal/dt/entity"
	"tung.gallery/pkg/models"
)

func (s *userService) UpdateUser(oldUser *entity.Users, req *user_dto.UserUpdateRequest) (*user_dto.UserUpdateResponse, error) {
	if req.Email != "" {
		_, err := s.Repo.ByEmail(req.Email)
		if err != nil && err != models.ErrNotFound {
			return nil, errors.New("error when check user email")
		} else if err == nil {
			return nil, errors.New("email has exist")
		}
	}

	user := &entity.Users{
		Email:    req.Email,
		Password: req.Password,
		Age:      req.Age,
		Birthday: req.Birthday,
		ImageURL: req.ImageURL,
	}

	err := s.Repo.Update(user, oldUser.ID)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("fail to update user profile")
	}

	return &user_dto.UserUpdateResponse{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
		Age:      user.Age,
		ImageURL: user.ImageURL,
		Birthday: user.Birthday,
	}, nil
}
