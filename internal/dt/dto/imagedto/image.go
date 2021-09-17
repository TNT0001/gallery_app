package imagedto

type (
	ImageUploadResponse struct {
	}

	ImageUploadRequest struct {
		UserID    int64  `json:"user_id"`
		GalleryID int64  `json:"gallery_id" ;binding:"required"`
		Title     string `json:"title" ;binding:"required"`
		ImageURL  string `json:"image_url"`
		ImageUUID string `json:"image_uuid"`
	}

	GetImageResponse struct {
		UserID    int64  `json:"user_id"`
		GalleryID int64  `json:"gallery_id"`
		Title     string `json:"title"`
		ImageURL  string `json:"image_url"`
		Content   string `json:"content"`
	}

	GetImageListResponse struct {
		Images []*GetImageResponse `json:"images"`
	}
)
