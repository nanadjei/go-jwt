package main

import (
	// "fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/nanadjei/go-jwt/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.DbConnect()
	initializers.Migration()
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
		"message": "pong",
		})
	})
	r.Run() 
}