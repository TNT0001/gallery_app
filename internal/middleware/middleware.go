package middleware

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"tung.gallery/internal/repo"
	"tung.gallery/pkg/models"
)

func AuthorizeJWT(repo repo.UserRepositoryInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, _ := c.Cookie("token")
		token, err := JWTAuthService().ValidateToken(tokenString)
		if !token.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		claims := token.Claims.(jwt.MapClaims)
		email := claims["email"].(string)
		user, err := repo.ByEmail(email)
		if err == models.ErrNotFound {
			c.AbortWithStatus(http.StatusBadRequest)
		} else if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
