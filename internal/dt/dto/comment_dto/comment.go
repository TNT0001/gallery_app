package comment_dto

type (
	CommentEditRequest struct {
		CommentID int64 `json:"comment_id"`
		Comment string `json:"comment"`
	}

	GetSingleCommentResponse struct {
		Comment string `json:"comment"`
		ImageID int64 `json:"image_id"`
		UserID int64 `json:"user_id"`
	}

	GetCommentByUserID struct {
		UserID int64 `json:"user_id"`
		UserName string `json:"user_name"`
		ListComment []*GetSingleCommentResponse `json:"list_comment"`
	}

	GetCommentByImageID struct {
		ImageID int64 `json:"image_id"`
		ListComment []*GetSingleCommentResponse `json:"list_comment"`
	}
)

