package user_repo

import "tung.gallery/internal/dt/entity"

type UserRepositoryInterface interface {
	CreateUser(user *entity.Users) error
	ByEmail(email string) (*entity.Users, error)
	ByID(id uint) (*entity.Users, error)
	Update(user *entity.Users, id uint) error
	Delete(id uint) error
	ByListID(id []uint) ([]*entity.Users, error)
}

type FriendRepositoryInterface interface {
	GetFriendIDList(id uint) ([]*entity.Friend, error)
	AddFriend(userID, FriendID uint) error
}
