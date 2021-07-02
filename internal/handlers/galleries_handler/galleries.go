package gale

import (
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
	c.HTML(http.StatusOK, "gallery", gin.H{
		"PageName": "contact",
	})
}

func (g *galleryHandler) Create(c *gin.Context) {
	req := dto.GalleryCreateRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		c.HTML(http.StatusBadRequest, "gallery", dto.GalleryCreateResponse{
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

	res, err := g.Services.CreateGallery(*user, req)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "gallery", res)
		return
	}

	url := fmt.Sprintf("/gallery/%d", res.ID)
	c.Request.Method = http.MethodGet
	c.Redirect(301, url)
}

func (g *galleryHandler) Show(c *gin.Context) {
	idString, ok := c.Params.Get("id")
	if !ok {
		c.HTML(http.StatusOK, "gallery", dto.ShowGalleryResponse{})
		return
	}

	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		c.HTML(http.StatusOK, "gallery", dto.ShowGalleryResponse{})
		return
	}

	res, err := g.Services.ShowGallery(uint(id))
	if err != nil {
		c.HTML(http.StatusOK, "gallery", res)
		return
	}

	c.HTML(http.StatusOK, "show", res)
}

func (g *galleryHandler) Edit(c *gin.Context) {
	idString, ok := c.Params.Get("id")
	if !ok {
		c.HTML(http.StatusOK, "gallery", dto.ShowGalleryResponse{})
		return
	}

	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		c.HTML(http.StatusOK, "gallery", dto.ShowGalleryResponse{})
		return
	}

	res, err := g.Services.ShowGallery(uint(id))
	if err != nil {
		c.HTML(http.StatusOK, "gallery", res)
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

	if res.UserID != user.ID {
		c.HTML(http.StatusForbidden, "home", dto.Alert{Level: services.AlertLvlInfo, Message: "Don't have permission"})
		return
	}

	c.HTML(http.StatusOK, "edit", res)
}
