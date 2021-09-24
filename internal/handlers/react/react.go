package react

import (
	"tung.gallery/internal/services/reactsservice"
)

type ReactHandler struct {
	Services reactsservice.ReactServiceInterface
}

func NewReactHandler(s reactsservice.ReactServiceInterface) *ReactHandler {
	return &ReactHandler{Services: s}
}
