package users

import (
	"tung.gallery/internal/dt/dto/user_dto"
	"tung.gallery/internal/dt/entity"
	"tung.gallery/internal/repo"
)

type UserServiceInterface interface {
	CreateUser(*user_dto.UserCreateRequest) (*user_dto.UserCreateResponse, error)
	UpdateUser(*entity.Users, *user_dto.UserUpdateRequest) (*user_dto.UserUpdateResponse, error)
	Login(*user_dto.UserLoginRequest) (*user_dto.UserLoginResponse, error)
	DeleteUser(*entity.Users) (*user_dto.UserDeleteResponse, error)
	FindUserById(uint) (*entity.Users, error)
	FindUserByEmail(*user_dto.UserLoginRequest) (*entity.Users, error)
}

type userService struct {
	Repo repo.UserRepositoryInterface
}

func NewUserService(r repo.UserRepositoryInterface) UserServiceInterface {
	return &userService{
		Repo: r,
	}
}

