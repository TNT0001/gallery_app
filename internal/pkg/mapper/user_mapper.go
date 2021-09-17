package mapper

import (
	"tung.gallery/internal/dt/dto/userdto"
	"tung.gallery/internal/dt/entity"
)

func FromUserToUserInfo(user *entity.Users) *userdto.UserGetInfoResponse {
	return &userdto.UserGetInfoResponse{
		Username: user.Username,
		Email:    user.Email,
		Age:      user.Age,
		ImageURL: user.ImageURL,
		Birthday: user.Birthday,
	}
}

func FromUserListToUserInfoList(users []*entity.Users) []*userdto.UserGetInfoResponse {
	res := make([]*userdto.UserGetInfoResponse, 0)
	for _, user := range users {
		res = append(res, FromUserToUserInfo(user))
	}
	return res
}
