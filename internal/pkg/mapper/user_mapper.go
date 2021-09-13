package mapper

import (
	"tung.gallery/internal/dt/dto/user_dto"
	"tung.gallery/internal/dt/entity"
)

func FromUserToUserInfo(user *entity.Users) *user_dto.UserGetInfoResponse {
	return &user_dto.UserGetInfoResponse{
		Username: user.Username,
		Email:    user.Email,
		Age:      user.Age,
		ImageURL: user.ImageURL,
		Birthday: user.Birthday,
	}
}

func FromUserListToUserInfoList(users []*entity.Users) []*user_dto.UserGetInfoResponse {
	res := make([]*user_dto.UserGetInfoResponse, 0)
	for _, user := range users {
		res = append(res, FromUserToUserInfo(user))
	}
	return res
}
