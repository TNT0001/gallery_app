package dto

type GalleryUpdateRequest struct {
	Title string `form:"title" binding:"required"`
	ID    uint
}
