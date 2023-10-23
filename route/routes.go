package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/randijulio13/gogin/middleware"
)

var Route *gin.Engine

func Setup() {
	Route.Use(middleware.Cors())
	Route.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
	})
	api()
}
