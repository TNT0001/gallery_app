package userservice

import (
	"errors"
	"log"
	"tung.gallery/internal/dt/dto/userdto"
	"tung.gallery/internal/pkg/mapper"
)

func (s *userService) GetFriendList(id int64) (*userdto.FriendListReponse, error) {
	friendList, err := s.Repo.GetFriendIDList(id)
	if err != nil {
		log.Println(err)
		return nil, errors.New("fail to get friend list")
	}

	friendIDList := make([]int64, 0)
	for _, friend := range friendList {
		friendIDList = append(friendIDList, int64(friend.FriendID))
	}

	FriendList, err := s.Repo.GetListUserByListID(friendIDList)
	if err != nil {
		log.Println(err)
		return nil, errors.New("fail to get friend info")
	}

	res := mapper.FromUserListToUserInfoList(FriendList)
	return &userdto.FriendListReponse{FriendList: res}, nil
}
