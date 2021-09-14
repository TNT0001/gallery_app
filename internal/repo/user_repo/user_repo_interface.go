package user_repo

import "tung.gallery/internal/dt/entity"

type UserRepositoryInterface interface {
	CreateUser(user *entity.Users) error
	ByEmail(email string) (*entity.Users, error)
	ByID(id int64) (*entity.Users, error)
	Update(user *entity.Users, id int64) error
	Delete(id int64) error
	ByListID(id []int64) ([]*entity.Users, error)
}

type FriendRepositoryInterface interface {
	GetFriendIDList(id int64) ([]*entity.Friend, error)
	AddFriend(userID, FriendID int64) error
}
