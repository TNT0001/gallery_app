package utils

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"tung.gallery/internal/dt/entity"
)

const (
	RememberTokenBytes = 32
)

var (
	// Error fail to get user from context
	ErrUserNotFound = errors.New("can't get user from context")

	// Error invalid file upload
	ErrInvalidFile = errors.New("file upload invalid")
)

func Bytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func String(nBytes int) (string, error) {
	b, err := Bytes(nBytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func RememberToken() (string, error) {
	return String(RememberTokenBytes)
}

// LoadDynamicTemplate func
// Return mutlitemplate render
func LoadDynamicTemplate(dir, baseNumber string) multitemplate.Renderer {
	r := multitemplate.New()

	base, err := filepath.Glob(dir + "/base/" + "*." + baseNumber + ".html")
	if err != nil {
		panic(err.Error())
	}

	layout, err := filepath.Glob(dir + "/layout/*.html")
	if err != nil {
		panic(err.Error())
	}

	pages, err := filepath.Glob(dir + "/page/*/*.html")
	if err != nil {
		panic(err.Error())
	}

	for _, page := range pages {
		files := base
		files = append(files, layout...)
		files = append(files, page)

		pageName := strings.TrimSuffix(filepath.Base(page), filepath.Ext(page))
		fmt.Println(pageName)
		r.AddFromFiles(pageName, files...)
	}
	return r
}

func GetUserFromContext(c *gin.Context) (*entity.Users, error) {
	u, exist := c.Get("user")
	if !exist {
		return nil, ErrUserNotFound
	}

	user, ok := u.(*entity.Users)
	if !ok {
		return nil, ErrUserNotFound
	}

	return user, nil
}

func CheckLogin(c *gin.Context) bool {
	_, err := GetUserFromContext(c)
	if err != nil {
		return false
	}
	return true
}

func CheckingExt(fileExt string, validExt []string) bool {
	for _, ext := range validExt {
		if fileExt == ext {
			return true
		}
	}
	return false
}
