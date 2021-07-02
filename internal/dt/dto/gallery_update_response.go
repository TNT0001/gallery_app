package dto

type GalleryUpdateResponse struct {
	ID    uint
	Title string `form:"title" binding:"required"`
	Login bool
	Alert Alert
}
