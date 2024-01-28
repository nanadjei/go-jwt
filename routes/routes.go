package routes

import (
	// "net/http"

	"github.com/gin-gonic/gin"

	"github.com/nanadjei/go-jwt/middleware"
	"github.com/nanadjei/go-jwt/controllers"
)

var Router = gin.Default()

func AppRoutes() {
	route := Router.Group("v1/")
	route.POST("/auth/signin", controllers.Signin)
	route.POST("/auth/signup", middleware.Guest, controllers.Signup)

	route.GET("/auth/user",  middleware.AuthCheck, controllers.AuthUser)
}