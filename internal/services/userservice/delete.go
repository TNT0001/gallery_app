package userservice

import (
	"errors"
	"log"
	"tung.gallery/internal/dt/dto/userdto"
	"tung.gallery/internal/dt/entity"
)

func (s *userService) DeleteUser(user *entity.Users) (*userdto.UserDeleteResponse, error) {
	err := s.Repo.DeleteUser(int64(user.ID))
	if err != nil {
		log.Println(err)
		return nil, errors.New("error when delete user")
	}

	return &userdto.UserDeleteResponse{}, nil
}
