package middleware

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"tung.gallery/internal/dt/dto"
	"tung.gallery/internal/repo"
	"tung.gallery/pkg/models"
	"tung.gallery/pkg/utils"
)

func AuthorizeJWT(repo repo.UserRepositoryInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, _ := c.Cookie("token")
		token, err := JWTAuthService().ValidateToken(tokenString)
		if err != nil {
			c.HTML(http.StatusUnauthorized, "login", dto.BaseResponse{})
			c.Abort()
			return
		}
		if !token.Valid {
			c.HTML(http.StatusUnauthorized, "login", dto.BaseResponse{})
			c.Abort()
			return
		}
		claims := token.Claims.(jwt.MapClaims)
		email := claims["email"].(string)
		user, err := repo.ByEmail(email)
		if err == models.ErrNotFound {
			c.HTML(http.StatusUnauthorized, "login", dto.BaseResponse{})
			c.Abort()
			return
		} else if err != nil {
			c.HTML(http.StatusUnauthorized, "login", dto.BaseResponse{})
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}

func LoginOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		login := utils.CheckLogin(c)
		if !login {
			c.Redirect(http.StatusFound, "/user/login")
			c.Abort()
			return
		}

		c.Next()
	}
}
