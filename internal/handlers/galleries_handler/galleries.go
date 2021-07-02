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
)

type galleryHandler struct {
	Services services.GalleriesServiceInterface
}

func NewGalleryHandler(s services.GalleriesServiceInterface) *galleryHandler {
	return &galleryHandler{Services: s}
}

func (g *galleryHandler) New(c *gin.Context) {
	c.HTML(http.StatusOK, "home", struct {
		Login bool
	}{Login: true})
}

func (g *galleryHandler) Create(c *gin.Context) {
	req := dto.GalleryCreateRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		c.HTML(http.StatusBadRequest, "home", dto.GalleryCreateResponse{
			Login: true,
			Alert: dto.Alert{
				Level:   services.AlertLvlInfo,
				Message: err.Error(),
			},
		})
		return
	}

	u, exists := c.Get("user")
	if !exists {
		c.Redirect(http.StatusTemporaryRedirect, "/user/login")
		return
	}

	user, ok := u.(*entity.Users)
	if !ok {
		c.Redirect(301, "/user/login")
		return
	}

	res, err := g.Services.CreateGallery(user, req)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "home", res)
		return
	}
	fmt.Println("\n", "\n", "\n", res, "\n", "\n")

	url := fmt.Sprintf("/gallery/%d", res.ID)
	c.Request.Method = http.MethodGet
	c.Redirect(301, url)
}

func (g *galleryHandler) Show(c *gin.Context) {
	idString, ok := c.Params.Get("id")
	if !ok {
		c.HTML(http.StatusBadRequest, "home", dto.ShowGalleryResponse{
			Login: true,
		})
		return
	}

	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		c.HTML(http.StatusBadRequest, "home", dto.ShowGalleryResponse{
			Login: true,
		})
		return
	}

	res, err := g.Services.ShowGallery(uint(id))
	if err != nil {
		c.HTML(http.StatusInternalServerError, "home", res)
		return
	}

	c.HTML(http.StatusOK, "show", res)
}

func (g *galleryHandler) Edit(c *gin.Context) {
	res, _, err := g.checkPermission(c)
	if err != nil {
		return
	}

	c.HTML(http.StatusOK, "edit", res)
}

func (g *galleryHandler) checkPermission(c *gin.Context) (dto.ShowGalleryResponse, *entity.Users, error) {
	idString, ok := c.Params.Get("id")
	if !ok {
		c.HTML(http.StatusBadRequest, "home", dto.ShowGalleryResponse{
			Login: true,
		})
		return dto.ShowGalleryResponse{Login: true}, nil, errors.New("not ok")
	}

	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		c.HTML(http.StatusBadRequest, "home", dto.ShowGalleryResponse{Login: true})
		return dto.ShowGalleryResponse{Login: true}, nil, err
	}

	res, err := g.Services.ShowGallery(uint(id))
	if err != nil {
		c.HTML(http.StatusInternalServerError, "home", res)
		return dto.ShowGalleryResponse{Login: true}, nil, err
	}

	u, exists := c.Get("user")
	if !exists {
		c.Redirect(http.StatusTemporaryRedirect, "/user/login")
		return dto.ShowGalleryResponse{
			Login: true,
		}, nil, errors.New("not ok")
	}

	user, ok := u.(*entity.Users)
	if !ok {
		c.Redirect(301, "/user/login")
		return dto.ShowGalleryResponse{
			Login: true,
		}, nil, errors.New("not ok")
	}

	if res.UserID != user.ID {
		c.HTML(http.StatusForbidden, "home", dto.Alert{Level: services.AlertLvlInfo, Message: "Don't have permission"})
		return dto.ShowGalleryResponse{
			Login: true,
		}, nil, err
	}
	return res, user, nil
}

func (g *galleryHandler) Update(c *gin.Context) {
	_, user, err := g.checkPermission(c)
	if err != nil {
		return
	}

	req := dto.GalleryUpdateRequest{}
	err = c.ShouldBind(&req)
	if err != nil {
		c.HTML(http.StatusOK, "edit", dto.ShowGalleryResponse{
			Login: true,
		})
		return
	}

	resUpdate, err := g.Services.Update(user, req)
	if err != nil {
		c.HTML(http.StatusForbidden, "edit", resUpdate)
		return
	}

	c.HTML(http.StatusOK, "edit", resUpdate)
}

func (g *galleryHandler) Delete(c *gin.Context) {
	res, _, err := g.checkPermission(c)
	if err != nil {
		return
	}

	resDelete, err := g.Services.Delete(res.ID)
	if err != nil {
		c.HTML(http.StatusForbidden, "home", resDelete)
		return
	}

	c.Redirect(http.StatusFound, "/gallery")
}

func (g *galleryHandler) ShowALlGalleries(c *gin.Context) {
	u, exists := c.Get("user")
	if !exists {
		c.Redirect(http.StatusTemporaryRedirect, "/user/login")
		return
	}

	user, ok := u.(*entity.Users)
	if !ok {
		c.Redirect(301, "/user/login")
		return
	}

	res, err := g.Services.GetAllGalleriesByUserID(user.ID)
	if err != nil {
		c.HTML(http.StatusForbidden, "home", res)
		return
	}

	c.HTML(http.StatusOK, "index", res)
}
