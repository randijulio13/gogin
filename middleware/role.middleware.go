package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/randijulio13/gogin/model"
)

func RoleMiddleware(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userContext, _ := c.Get("user")
		user, _ := userContext.(*model.User)

		if can := user.Can(role); !can {
			c.JSON(http.StatusForbidden, gin.H{"message": "you shall not pass!"})
			c.Abort()
			return
		}
		c.Next()
	}
}
