package dto

type BaseResponse struct {
	Alert
}

type Alert struct {
	Level   string
	Message string
}
