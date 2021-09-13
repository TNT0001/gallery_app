package users

import (
	"tung.gallery/internal/dt/dto/user_dto"
	"tung.gallery/internal/dt/entity"
	"tung.gallery/internal/repo/user_repo"
)

type UserServiceInterface interface {
	GetFriendList(id uint) (*user_dto.FriendListReponse, error)
	CreateUser(*user_dto.UserCreateRequest) (*user_dto.UserCreateResponse, error)
	UpdateUser(*entity.Users, *user_dto.UserUpdateRequest) (*user_dto.UserUpdateResponse, error)
	Login(*user_dto.UserLoginRequest) (*user_dto.UserLoginResponse, error)
	DeleteUser(*entity.Users) (*user_dto.UserDeleteResponse, error)
	FindUserById(uint) (*entity.Users, error)
	FindUserByEmail(*user_dto.UserLoginRequest) (*entity.Users, error)
	AddFriend(userID uint, req *user_dto.AddFriendRequest) error
}

type userService struct {
	UserRepo user_repo.UserRepositoryInterface
	FriendRepo user_repo.FriendRepositoryInterface
}

func NewUserService(ur user_repo.UserRepositoryInterface, fr user_repo.FriendRepositoryInterface) UserServiceInterface {
	return &userService{
		UserRepo: ur,
		FriendRepo: fr,
	}
}

