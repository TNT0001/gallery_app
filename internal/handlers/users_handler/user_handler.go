package users

import (
	"tung.gallery/internal/services/userservice"
)

type userHandler struct {
	Services userservice.UserServiceInterface
}

func NewUserHandler(s userservice.UserServiceInterface) *userHandler {
	return &userHandler{Services: s}
}
