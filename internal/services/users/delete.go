package users

import (
	"tung.gallery/internal/dt/dto/user_dto"
	"tung.gallery/internal/dt/entity"
)

func (s *userService) DeleteUser(user *entity.Users) (*user_dto.UserDeleteResponse, error) {
	err := s.Repo.Delete(user.ID)
	if err != nil {
		return nil, err
	}

	return &user_dto.UserDeleteResponse{}, nil
}
