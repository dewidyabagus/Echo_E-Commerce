package api

import (
	echo "github.com/labstack/echo/v4"

	"RESTful/api/middleware"
	"RESTful/api/v1/admin"
	"RESTful/api/v1/auth"
	"RESTful/api/v1/user"
)

type Routing struct {
	User  *user.Controller
	Auth  *auth.Controller
	Admin *admin.Controller
}

func RegisterRouters(e *echo.Echo, routing *Routing) {
	user := routing.User
	auth := routing.Auth
	admin := routing.Admin

	// IAM service
	authGroup := e.Group("/v1/auth")
	authGroup.POST("/users", auth.LoginUserWithEmailPassword)
	authGroup.POST("/admins", auth.LoginAdminWithEmailPassword)

	// Register Admin
	e.POST("/v1/admins/register", admin.AddNewAdmin)

	// User Information Changes
	userGroup := e.Group("/v1/users")
	userGroup.Use(middleware.JWTMiddleware())
	userGroup.GET("", user.FindAllUser)
	userGroup.POST("/register", user.AddNewUser)
	userGroup.GET("/:id", user.FindUserByUserId)
	userGroup.PUT("/:id", user.UpdateUser)
	userGroup.DELETE("/:id", user.DeleteUser)
}
