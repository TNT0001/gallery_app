package dto

type ShowGalleryResponse struct {
	Title  string
	ID     uint
	UserID uint
	Login  bool
	Alert  Alert
}
