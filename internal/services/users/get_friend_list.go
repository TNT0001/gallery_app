package users

import (
	"errors"
	"log"
	"tung.gallery/internal/dt/dto/user_dto"
	"tung.gallery/internal/pkg/mapper"
)

func (s *userService) GetFriendList(id uint) (*user_dto.FriendListReponse, error) {
	friendList, err := s.FriendRepo.GetFriendIDList(id)
	if err != nil {
		log.Println(err)
		return nil, errors.New("fail to get friend list")
	}

	friendIDList := make([]uint, 0)
	for _, friend := range friendList {
		friendIDList = append(friendIDList, friend.FriendID)
	}

	FriendList, err := s.UserRepo.ByListID(friendIDList)
	if err != nil {
		log.Println(err)
		return nil, errors.New("fail to get friend info")
	}

	res := mapper.FromUserListToUserInfoList(FriendList)
	return &user_dto.FriendListReponse{FriendList: res}, nil
}
