package gallery_dto

type ShowGalleryResponse struct {
	Title  string `json:"title"`
	ID     uint `json:"id"`
	UserID uint `json:"user_id"`
	ImageUrls []string `json:"image_urls"`
}
