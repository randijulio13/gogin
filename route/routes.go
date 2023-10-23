package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var Route *gin.Engine

func Setup() {
	Route.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
	})
	api()
}
