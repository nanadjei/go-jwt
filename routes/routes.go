package routes

import (
	// "net/http"

	"github.com/gin-gonic/gin"

	// "github.com/nanadjei/go-jwt/middleware"
	"github.com/nanadjei/go-jwt/controllers"
)

var Router = gin.Default()

func AppRoutes() {
	route := Router.Group("v1/")
	route.POST("/auth/signin", controllers.Signin)
	route.POST("/auth/signup", controllers.Signup)
	route.GET("/auth/signout", controllers.Signout)

	route.GET("/auth/user", controllers.AuthUser)

	route.POST("/password/reset", controllers.ForgotPassword)
	route.POST("/password/verify", controllers.VerifyOTPcode)

	// route.GET("/tester", controllers.PrepareOTPcode)
}