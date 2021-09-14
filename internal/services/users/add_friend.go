package users

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"tung.gallery/internal/dt/dto/user_dto"
	"tung.gallery/internal/dt/entity"
)

func (s *userService) AddFriend(userID int64, req *user_dto.AddFriendRequest) error {
	var err error
	var user *entity.Users
	if req.ID != 0 {
		user, err = s.UserRepo.ByID(int64(req.ID))
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

	friendList, err := s.FriendRepo.GetFriendIDList(userID)
	if err != nil {
		log.Println(err.Error())
		return errors.New("error when check friend of user")
	}

	for _, f := range friendList {
		if f.FriendID == req.ID {
			return errors.New("you have already is friend of that user")
		}
	}

	err = s.FriendRepo.AddFriend(userID, int64(user.ID))
	if err != nil {
		log.Println(err)
		return errors.New("fail to add friend")
	}
	return nil
}
