package main

import (
	"github.com/nanadjei/go-jwt/initializers"
	"github.com/nanadjei/go-jwt/routes"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.DbConnect()
	initializers.Migration()
}

func main() {
	routes.AppRoutes()
	routes.Router.Run() 
}