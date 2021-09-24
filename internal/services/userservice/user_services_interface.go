package userservice

import (
	"tung.gallery/internal/dt/dto/userdto"
	"tung.gallery/internal/dt/entity"
	"tung.gallery/internal/repo"
)

type UserServiceInterface interface {
	GetFriendList(id int64) (*userdto.FriendListReponse, error)
	CreateUser(*userdto.UserCreateRequest) (*userdto.UserCreateResponse, error)
	UpdateUser(*entity.Users, *userdto.UserUpdateRequest) (*userdto.UserUpdateResponse, error)
	Login(*userdto.UserLoginRequest) (*userdto.UserLoginResponse, error)
	DeleteUser(*entity.Users) (*userdto.UserDeleteResponse, error)
	FindUserByID(int64) (*entity.Users, error)
	FindUserByEmail(*userdto.UserLoginRequest) (*entity.Users, error)
	AddFriend(userID int64, req *userdto.AddFriendRequest) error
}

type userService struct {
	Repo repo.Repo
}

func NewUserService(r repo.Repo) UserServiceInterface {
	return &userService{
		r,
	}
}
