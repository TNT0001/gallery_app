package dto

type Gallrery struct {
	Title  string `json:"title"`
	ID     uint `json:"id"`
	UserID uint `json:"user_id"`
}

type UserGetAllGalleriesResponse struct {
	Galleries []Gallrery `json:"galleries"`
}
