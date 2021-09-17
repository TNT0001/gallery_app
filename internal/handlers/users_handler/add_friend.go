package users

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"tung.gallery/internal/dt/dto/userdto"
	"tung.gallery/internal/pkg"
	"tung.gallery/pkg/utils"
)

func (u *userHandler) AddFriend(c *gin.Context) {
	user, err := utils.GetUserFromContext(c)
	if err != nil {
		log.Println(err.Error())
		pkg.ResponseErrorJSON(c, http.StatusBadRequest, err.Error())
		return
	}

	req := &userdto.AddFriendRequest{}
	err = c.BindJSON(req)
	if err != nil {
		log.Println(err)
		pkg.ResponseErrorJSON(c, http.StatusBadRequest, "fail to parse request")
		return
	}

	err = valid(req, user.ID)
	if err != nil {
		log.Println(err.Error())
		pkg.ResponseErrorJSON(c, http.StatusBadRequest, err.Error())
		return
	}

	err = u.Services.AddFriend(int64(user.ID), req)
	if err != nil {
		pkg.ResponseErrorJSON(c, http.StatusInternalServerError, err.Error())
		return
	}

	pkg.ResponseSuccessJSON(c, http.StatusOK, nil)
}

func valid(req *userdto.AddFriendRequest, userID uint) error {
	if req.Email == "" && req.ID < 1 {
		return errors.New("need at least one of email or id of user to add friend")
	}
	if req.ID == userID {
		return errors.New("can't add friend to yourself")
	}
	return nil
}
