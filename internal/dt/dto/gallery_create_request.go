package dto

type GalleryCreateRequest struct {
	Title string `form:"title" binding:"required"`
}
