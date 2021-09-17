package repo

import (
	"gorm.io/gorm"
	"tung.gallery/internal/repo/commentrepo"
	"tung.gallery/internal/repo/galleryrepo"
	"tung.gallery/internal/repo/imagerepo"
	"tung.gallery/internal/repo/reactrepo"
	"tung.gallery/internal/repo/userrepo"
	"tung.gallery/internal/repo/userrepo/friend"
	"tung.gallery/internal/repo/userrepo/user"
)

type GalleryRepository interface {
	userrepo.UserRepositoryInterface
	userrepo.FriendRepositoryInterface
	commentrepo.CommentRepoInterface
	reactrepo.ReactRepositoryInterface
	galleryrepo.GalleryRepositoryInterface
	imagerepo.ImageReposirotyInterface
}

type repo struct {
	userrepo.UserRepositoryInterface
	userrepo.FriendRepositoryInterface
	commentrepo.CommentRepoInterface
	reactrepo.ReactRepositoryInterface
	galleryrepo.GalleryRepositoryInterface
	imagerepo.ImageReposirotyInterface
}

func NewRepo(DB *gorm.DB) GalleryRepository {
	return &repo{
		UserRepositoryInterface:    user.NewUserRepo(DB),
		FriendRepositoryInterface:  friend.NewFriendRepo(DB),
		CommentRepoInterface:       commentrepo.NewCommentRepo(DB),
		ReactRepositoryInterface:   reactrepo.NewReactRepo(DB),
		GalleryRepositoryInterface: galleryrepo.NewGalleryRepo(DB),
		ImageReposirotyInterface:   imagerepo.NewImageRepository(DB),
	}
}
