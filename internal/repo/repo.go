package repo

import (
	"fmt"
	"gorm.io/gorm"
	"tung.gallery/internal/repo/commentrepo"
	"tung.gallery/internal/repo/galleryrepo"
	"tung.gallery/internal/repo/imagerepo"
	"tung.gallery/internal/repo/reactrepo"
	"tung.gallery/internal/repo/userrepo"
	"tung.gallery/internal/repo/userrepo/friend"
	"tung.gallery/internal/repo/userrepo/user"
)

type Repo interface {
	userrepo.UserRepositoryInterface
	userrepo.FriendRepositoryInterface
	commentrepo.CommentRepoInterface
	reactrepo.ReactRepositoryInterface
	galleryrepo.GalleryRepositoryInterface
	imagerepo.ImageReposirotyInterface
	Transactions(f func(subRepo Repo) error) (err error)
}

type repo struct {
	userrepo.UserRepositoryInterface
	userrepo.FriendRepositoryInterface
	commentrepo.CommentRepoInterface
	reactrepo.ReactRepositoryInterface
	galleryrepo.GalleryRepositoryInterface
	imagerepo.ImageReposirotyInterface
	*gorm.DB
}

func NewRepo(DB *gorm.DB) Repo {
	return &repo{
		UserRepositoryInterface:    user.NewUserRepo(DB),
		FriendRepositoryInterface:  friend.NewFriendRepo(DB),
		CommentRepoInterface:       commentrepo.NewCommentRepo(DB),
		ReactRepositoryInterface:   reactrepo.NewReactRepo(DB),
		GalleryRepositoryInterface: galleryrepo.NewGalleryRepo(DB),
		ImageReposirotyInterface:   imagerepo.NewImageRepository(DB),
		DB:                         DB,
	}
}

func (r *repo) Transactions(f func(subRepo Repo) error) (err error) {
	subDB := r.DB.Begin()
	subRepo := NewRepo(subDB)

	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("panic %v", p)
			if rErr := subDB.Rollback().Error; rErr != nil {
				err = fmt.Errorf("%v, %v", err, rErr)
			}
		}
	}()

	err = f(subRepo)
	if err != nil {
		if rErr := subDB.Rollback().Error; rErr != nil {
			err = fmt.Errorf("%v, %v", err, rErr)
			return
		}
	}

	return subDB.Commit().Error
}
