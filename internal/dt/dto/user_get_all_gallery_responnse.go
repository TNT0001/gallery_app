package dto

type Gallrery struct {
	Title  string
	ID     uint
	UserID uint
}

type UserGetAllGalleriesResponse struct {
	Galleries []Gallrery
	BaseResponse
}
