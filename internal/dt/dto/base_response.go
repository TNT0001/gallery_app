package dto

type (
	BaseResponse struct {
		Success bool `json:"success"`
		DataMsg
		ErrorMsg
	}

	DataMsg struct {
		Data interface{} `json:"data,omitempty"`
	}

	ErrorMsg struct {
		Error interface{} `json:"error,omitempty"`
	}
)
