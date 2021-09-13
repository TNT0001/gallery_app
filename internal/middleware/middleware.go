package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"tung.gallery/internal/repo/user_repo"
	"tung.gallery/pkg/models"
	"tung.gallery/pkg/utils"
)

func AuthorizeJWT(repo user_repo.UserRepositoryInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		token, err := JWTAuthService().ValidateToken(tokenString[7:])
		if err != nil {
			return
		}
		if !token.Valid {
			return
		}
		claims := token.Claims.(jwt.MapClaims)
		email := claims["email"].(string)
		user, err := repo.ByEmail(email)
		if err == models.ErrNotFound {
			return
		} else if err != nil {
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
