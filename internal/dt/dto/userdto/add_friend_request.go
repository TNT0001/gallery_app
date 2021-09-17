package userdto

type AddFriendRequest struct {
	Email string `json:"email" binding:"email"`
	ID    uint   `json:"id"`
}
