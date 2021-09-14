package comment

import (
	"tung.gallery/internal/repo/comment_repo"
	"tung.gallery/internal/repo/gallery_repo"
	"tung.gallery/internal/repo/image_repo"
	"tung.gallery/internal/repo/user_repo"
)

type CommentServiceInterface interface {
}

type commentService struct {
	CommentRepo comment_repo.CommentRepoInterface
	GalleryRepo gallery_repo.GalleryRepositoryInterface
	ImageRepo image_repo.ImageReposirotyInterface
	UserRepo user_repo.UserRepositoryInterface
}

func NewCommentService(cr comment_repo.CommentRepoInterface, gr gallery_repo.GalleryRepositoryInterface,
	ir image_repo.ImageReposirotyInterface, ur user_repo.UserRepositoryInterface) *commentService {
	return &commentService{
		cr,
		gr,
		ir,
		ur,
	}
}

