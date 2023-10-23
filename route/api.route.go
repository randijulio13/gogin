package route

import (
	"github.com/gin-contrib/cors"
	"github.com/randijulio13/gogin/handler"
	"github.com/randijulio13/gogin/middleware"
)

func api() {
	Route.Use(cors.Default())

	// config := cors.DefaultConfig()
	// // config.AllowOrigins = []string{"http://google.com"}
	// // config.AllowOrigins = []string{"http://google.com", "http://facebook.com"}
	// config.AllowAllOrigins = true
	// config.AllowHeaders = []string{"Token", "Origin", "Access-Control-Allow-Origin"}

	// Route.Use(cors.New(config))

	api := Route.Group("/api")
	api.GET("/", handler.GetAllUser)
	api.POST("/login", handler.Login)

	user := Route.Group("/api/user")
	user.Use(middleware.AuthMiddleware).Use(middleware.RoleMiddleware("user_management"))

	user.GET("/", handler.GetAllUser)
	user.GET("/:id", handler.GetUser)
	user.POST("/", handler.CreateUser)
	user.PATCH("/:id", handler.EditUser)
}
