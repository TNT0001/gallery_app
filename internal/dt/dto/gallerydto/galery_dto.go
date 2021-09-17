package gallerydto

type GalleryCreateRequest struct {
	Title string `json:"title" binding:"required"`
}

type GalleryCreateResponse struct {
	Title  string `json:"title"`
	ID     uint   `json:"id"`
	UserID uint   `json:"user_id"`
}

type GalleryDeleteResponse struct {
}

type ShowGalleryResponse struct {
	Title  string `json:"title"`
	ID     int64  `json:"id"`
	UserID int64  `json:"user_id"`
}

type ShowAllGalleryByUserIDResponse struct {
	Galleries []*ShowGalleryResponse
}

type GalleryUpdateRequest struct {
	Title string `json:"title" binding:"required"`
	ID    int64  `json:"id"`
}

type GalleryUpdateResponse struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
}

type GalleryEditResponse struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
}
