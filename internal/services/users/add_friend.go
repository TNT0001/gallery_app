package users

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"tung.gallery/internal/dt/dto/user_dto"
	"tung.gallery/internal/dt/entity"
)

func (s *userService) AddFriend(userID uint, req *user_dto.AddFriendRequest) error {
	var err error
	var user *entity.Users
	if req.ID != 0 {
		user, err = s.UserRepo.ByID(req.ID)
	} else {
		user, err = s.UserRepo.ByEmail(req.Email)
	}

	if err != nil {
		log.Println(err.Error())
		if err != gorm.ErrRecordNotFound {
			return errors.New("fail to get user info by id")
		}
		return errors.New("user don't exists")
	}

	err = s.FriendRepo.AddFriend(userID, user.ID)
	if err != nil {
		log.Println(err)
		return errors.New("fail to add friend")
	}
	return nil
}
