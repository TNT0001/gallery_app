package gale

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"tung.gallery/internal/dt/dto"
	"tung.gallery/internal/dt/entity"
	"tung.gallery/internal/services"
	"tung.gallery/pkg/models"
	"tung.gallery/pkg/utils"
)

var (
	// Error invalid request
	ErrInvalidRequest = errors.New("invalid request")

	// Error Permisson request fail
	ErrInvalidPermission = errors.New("invalid permission")
)

type galleryHandler struct {
	Services services.GalleriesServiceInterface
}

func NewGalleryHandler(s services.GalleriesServiceInterface) *galleryHandler {
	return &galleryHandler{Services: s}
}

func (g *galleryHandler) GetNewGalleryPage(c *gin.Context) {
	login := utils.CheckLogin(c)
	if !login {
		c.Redirect(http.StatusFound, "/user/login")
		return
	}

	baseResponse := dto.BaseResponse{Login: login}
	c.HTML(http.StatusOK, "new_gallery", baseResponse)
}

func (g *galleryHandler) NewGallery(c *gin.Context) {
	login := utils.CheckLogin(c)
	if !login {
		c.Redirect(http.StatusFound, "/user/login")
		return
	}

	req := dto.GalleryCreateRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		baseResponse := utils.BaseResponse(login, services.AlertLvlInfo, ErrInvalidRequest.Error())
		c.HTML(http.StatusBadRequest, "new_gallery", dto.GalleryCreateResponse{
			BaseResponse: baseResponse,
		})
		return
	}

	user, err := utils.GetUserFromContext(c)
	if err != nil {
		c.Redirect(http.StatusFound, "/user/login")
		return
	}

	res, err := g.Services.CreateGallery(user, req)
	res.Login = login
	if err != nil {
		c.HTML(http.StatusInternalServerError, "new_gallery", res)
		return
	}

	url := fmt.Sprintf("/gallery/%d", res.ID)
	c.Redirect(http.StatusFound, url)
}

func (g *galleryHandler) GetGalleryPage(c *gin.Context) {
	login := utils.CheckLogin(c)
	if !login {
		c.Redirect(http.StatusFound, "/user/login")
		return
	}

	idString, ok := c.Params.Get("id")
	if !ok {
		c.Redirect(http.StatusFound, "all_galleries")
		return
	}

	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		c.Redirect(http.StatusFound, "all_galleries")
		return
	}

	res, err := g.Services.ShowGallery(uint(id))
	res.Login = login
	if err != nil {
		c.Redirect(http.StatusFound, "all_galleries")
		return
	}

	c.HTML(http.StatusOK, "show_gallery", res)
}

func (g *galleryHandler) GetEditPage(c *gin.Context) {
	login := utils.CheckLogin(c)
	if !login {
		c.Redirect(http.StatusFound, "/user/login")
		return
	}

	baseResponse, id, _, err := g.checkPermission(c)
	baseResponse.Login = login
	if err != nil {
		c.HTML(http.StatusNonAuthoritativeInfo, "home", baseResponse)
		return
	}

	gallery, err := g.Services.ShowGallery(uint(id))
	if err != nil {
		c.HTML(http.StatusNonAuthoritativeInfo, "home", baseResponse)
		return
	}

	res := dto.GalleryEditResponse{Title: gallery.Title, ID: gallery.ID, BaseResponse: baseResponse}
	c.HTML(http.StatusOK, "edit_gallery", res)
}

func (g *galleryHandler) EditGallery(c *gin.Context) {
	login := utils.CheckLogin(c)
	if !login {
		c.Redirect(http.StatusFound, "/user/login")
		return
	}

	res, _, user, err := g.checkPermission(c)
	if err != nil {
		c.HTML(http.StatusNonAuthoritativeInfo, "home", res)
		return
	}

	req := dto.GalleryUpdateRequest{}
	err = c.ShouldBind(&req)
	if err != nil {
		baseResponse := utils.BaseResponse(login, services.AlertLvlError, ErrInvalidRequest.Error())
		c.HTML(http.StatusBadRequest, "edit_gallery", baseResponse)
		return
	}

	resUpdate, err := g.Services.Update(user, req)
	resUpdate.Login = login
	if err != nil {
		c.HTML(http.StatusInternalServerError, "edit_gallery", resUpdate)
		return
	}

	c.HTML(http.StatusOK, "edit_gallery", resUpdate)
}

func (g *galleryHandler) Delete(c *gin.Context) {
	login := utils.CheckLogin(c)
	if !login {
		c.Redirect(http.StatusFound, "/user/login")
		return
	}

	res, id, _, err := g.checkPermission(c)
	if err != nil {
		c.HTML(http.StatusNonAuthoritativeInfo, "home", res)
		return
	}

	resDelete, err := g.Services.Delete(uint(id))
	resDelete.Login = login
	if err != nil {
		c.HTML(http.StatusInternalServerError, "edit_gallery", resDelete)
		return
	}

	c.Redirect(http.StatusFound, "/gallery")
}

func (g *galleryHandler) ShowALlGalleries(c *gin.Context) {
	login := utils.CheckLogin(c)
	if !login {
		c.Redirect(http.StatusFound, "/user/login")
		return
	}

	user, err := utils.GetUserFromContext(c)
	if err != nil {
		c.Redirect(http.StatusFound, "/user/login")
		return
	}

	res, err := g.Services.GetAllGalleriesByUserID(user.ID)
	res.Login = login
	if err != nil {
		c.HTML(http.StatusInternalServerError, "home", res.BaseResponse)
		return
	}

	c.HTML(http.StatusOK, "all_galleries", res)
}

func (g *galleryHandler) checkPermission(c *gin.Context) (dto.BaseResponse, int, *entity.Users, error) {
	idString, ok := c.Params.Get("id")
	if !ok {
		baseResponse := utils.BaseResponse(false, services.AlertLvlError, ErrInvalidRequest.Error())
		return baseResponse, -1, nil, ErrInvalidRequest
	}

	id, err := strconv.Atoi(idString)
	if err != nil {
		baseResponse := utils.BaseResponse(false, services.AlertLvlError, ErrInvalidRequest.Error())
		return baseResponse, -1, nil, ErrInvalidRequest
	}

	gallery, err := g.Services.ShowGallery(uint(id))
	if err != nil {
		baseResponse := utils.BaseResponse(false, services.AlertLvlError, models.ErrInternalServerError.Error())
		return baseResponse, -1, nil, models.ErrInternalServerError
	}

	user, err := utils.GetUserFromContext(c)
	if err != nil {
		baseResponse := utils.BaseResponse(false, services.AlertLvlError, utils.ErrUserNotFound.Error())
		return baseResponse, -1, nil, utils.ErrUserNotFound
	}

	if gallery.UserID != user.ID {
		baseResponse := utils.BaseResponse(false, services.AlertLvlError, ErrInvalidPermission.Error())
		return baseResponse, -1, nil, utils.ErrUserNotFound
	}

	baseResponse := dto.BaseResponse{}
	return baseResponse, id, user, nil
}
