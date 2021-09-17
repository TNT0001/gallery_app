package userdto

type FriendListReponse struct {
	FriendList []*UserGetInfoResponse `json:"friend_list"`
}
