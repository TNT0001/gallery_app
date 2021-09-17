package userrepo

import "tung.gallery/internal/dt/entity"

type UserRepositoryInterface interface {
	CreateUser(user *entity.Users) error
	GetUserByEmail(email string) (*entity.Users, error)
	GetUserByID(id int64) (*entity.Users, error)
	UpdateUser(user *entity.Users, id int64) error
	DeleteUser(id int64) error
	GetListUserByListID(id []int64) ([]*entity.Users, error)
}

type FriendRepositoryInterface interface {
	GetFriendIDList(id int64) ([]*entity.Friend, error)
	AddFriend(userID, FriendID int64) error
}
