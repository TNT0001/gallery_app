package userservice

import (
	"tung.gallery/internal/dt/dto/userdto"
	"tung.gallery/internal/middleware"
)

func (s *userService) Login(req *userdto.UserLoginRequest) (*userdto.UserLoginResponse, error) {
	user, err := s.FindUserByEmail(req)
	if err != nil {
		return nil, err
	}

	token := middleware.JWTAuthService().GenerateToken(user.Email, true)

	return &userdto.UserLoginResponse{Token: token}, nil
}
