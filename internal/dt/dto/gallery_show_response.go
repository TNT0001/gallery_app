package dto

type ShowGalleryResponse struct {
	Title  string
	ID     uint
	UserID uint
	Images []string
	BaseResponse
}
