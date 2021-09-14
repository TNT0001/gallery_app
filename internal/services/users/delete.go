package users

import (
	"errors"
	"log"
	"tung.gallery/internal/dt/dto/user_dto"
	"tung.gallery/internal/dt/entity"
)

func (s *userService) DeleteUser(user *entity.Users) (*user_dto.UserDeleteResponse, error) {
	err := s.UserRepo.Delete(int64(user.ID))
	if err != nil {
		log.Println(err)
		return nil, errors.New("error when delete user")
	}

	return &user_dto.UserDeleteResponse{}, nil
}
