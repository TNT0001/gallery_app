package dto

type GalleryCreateResponse struct {
	Title  string `form:"title"`
	ID     uint
	UserID uint
	Login  bool
	Alert  Alert
}
