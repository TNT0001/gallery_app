package gallery_dto

type GalleryCreateRequest struct {
	Title string `json:"title" binding:"required"`
}
