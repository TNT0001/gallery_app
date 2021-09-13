package users

import (
	"tung.gallery/internal/dt/dto/user_dto"
	"tung.gallery/internal/middleware"
)

func (s *userService) Login(req *user_dto.UserLoginRequest) (*user_dto.UserLoginResponse, error) {
	user, err := s.FindUserByEmail(req)
	if err != nil {
		return nil, err
	}

	token := middleware.JWTAuthService().GenerateToken(user.Email, true)

	return &user_dto.UserLoginResponse{Token: token}, nil
}
