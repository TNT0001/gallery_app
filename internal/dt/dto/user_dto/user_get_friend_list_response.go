package user_dto

type FriendListReponse struct {
	FriendList []*UserGetInfoResponse `json:"friend_list"`
}
