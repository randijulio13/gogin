package route

import (
	"github.com/gin-contrib/cors"
	"github.com/randijulio13/gogin/handler"
	"github.com/randijulio13/gogin/middleware"
)

func api() {
	// Route.Use(cors.Default())

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = append(config.AllowHeaders, "Access-Control-Allow-Origin")
	config.AllowHeaders = append(config.AllowHeaders, "Token")
	Route.Use(cors.New(config))

	api := Route.Group("api")
	api.POST("login", handler.Login)

	user := api.Group("user")
	user.Use(middleware.AuthMiddleware).Use(middleware.RoleMiddleware("user_management"))

	user.GET("", handler.GetAllUser)
	user.GET(":id", handler.GetUser)
	user.POST("", handler.CreateUser)
	user.PATCH(":id", handler.EditUser)
}
