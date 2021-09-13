package gallery_dto

type GalleryCreateResponse struct {
	Title  string `json:"title"`
	ID     uint `json:"id"`
	UserID uint `json:"user_id"`
	Login  bool `json:"login"`
}
