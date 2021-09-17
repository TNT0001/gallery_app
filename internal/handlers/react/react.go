package react

import (
	"tung.gallery/internal/services/reactsservice"
)

type reactHandler struct {
	Services reactsservice.ReactServiceInterface
}

func NewReactHandler(s reactsservice.ReactServiceInterface) *reactHandler {
	return &reactHandler{Services: s}
}
