package gallery_dto

type GalleryUpdateRequest struct {
	Title string `json:"title" binding:"required"`
	ID    uint `json:"id"`
}
