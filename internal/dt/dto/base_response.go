package dto

type BaseResponse struct {
	Login bool
	Alert Alert
}

type Alert struct {
	Level   string
	Message string
}
