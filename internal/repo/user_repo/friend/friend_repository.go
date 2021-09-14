package friend

import (
	"gorm.io/gorm"
	"tung.gallery/internal/dt/entity"
	"tung.gallery/internal/repo/user_repo"
)

type friendRepo struct {
	DB *gorm.DB
}

func NewFriendRepo(db *gorm.DB) user_repo.FriendRepositoryInterface {
	return &friendRepo{DB: db}
}

func (r *friendRepo) GetFriendIDList(id int64) ([]*entity.Friend, error) {
	friendList := make([]*entity.Friend, 0)
	err := r.DB.Where("user_id = ?", id).Find(&friendList).Error
	if err != nil {
		return nil, err
	}

	return friendList, nil
}

func (r *friendRepo) AddFriend(userID, FriendID int64) error {
	friend := []*entity.Friend{
		{
			UserID:   uint(userID),
			FriendID: uint(FriendID),
		},
		{
			UserID: uint(FriendID),
			FriendID: uint(userID),
		},
	}
	return r.DB.Create(&friend).Error
}
