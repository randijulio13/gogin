package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/randijulio13/gogin/model"
)

func validateToken(tokenString string) (bool, jwt.MapClaims) {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if tokenString == "" {
		return false, nil
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return false, nil
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid && err != nil {
		return false, nil
	}

	return true, claims
}

func AuthMiddleware(c *gin.Context) {
	tokenString := c.GetHeader("Token")
	valid, claims := validateToken(tokenString)

	if !valid {
		c.JSON(http.StatusForbidden, gin.H{"message": "invalid token"})
		c.Abort()
		return
	}

	nip := claims["nip"].(string)
	var user model.User
	user.GetByNip(nip)

	c.Set("user", &user)
	c.Next()
}
