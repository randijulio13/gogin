package route

import (
	"github.com/randijulio13/gogin/handler"
	"github.com/randijulio13/gogin/middleware"
)

func api() {
	api := Route.Group("api")
	api.POST("login", handler.Login)

	user := api.Group("user")
	user.Use(middleware.AuthMiddleware).Use(middleware.RoleMiddleware("user_management"))

	user.GET("", handler.GetAllUser)
	user.GET(":id", handler.GetUser)
	user.POST("", handler.CreateUser)
	user.PATCH(":id", handler.EditUser)
}
