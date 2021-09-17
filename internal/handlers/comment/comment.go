package comment

import "tung.gallery/internal/services/commentservice"

type commentHandler struct {
	CommentService commentservice.CommentServiceInterface
}

func NewCommentHandler(cs commentservice.CommentServiceInterface) *commentHandler {
	return &commentHandler{
		CommentService: cs,
	}
}
