package reactdto

type (
	CreateReactRequest struct {
		TypeID  int64 `json:"type_id,omitempty"`
		ImageID int64 `json:"image_id,omitempty"`
		UserID  int64 `json:"user_id,omitempty"`
	}

	TotalReactIDCountByType struct {
		TypeID int64 `json:"type_id"`
		Total  int64 `json:"total"`
	}

	TotalReactCountByType struct {
		Type  string `json:"type_id"`
		Total int64  `json:"total"`
	}

	GetReactByIDResponse struct {
		Type    string `json:"type_id"`
		UserID  int64  `json:"user_id"`
		ImageID int64  `json:"image_id"`
	}

	GetReactTotalByImageIDResponse struct {
		ImageID int64 `json:"image_id"`
		Total   int64 `json:"total"`
	}

	GetReactTotalByEachTypeAndImageIDResponse struct {
		TotalForEachType []*TotalReactCountByType `json:"total_for_each_type"`
	}

	GetListReactByUserIDResponse struct {
		UserID int64                   `json:"user_id"`
		Reacts []*GetReactByIDResponse `json:"reacts"`
	}

	GetListReactByImagesIDResponse struct {
		ImageID int64                   `json:"image_id"`
		Reacts  []*GetReactByIDResponse `json:"reacts"`
	}
)
