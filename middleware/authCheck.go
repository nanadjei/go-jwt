package middleware

import (
	"os"
	"net/http"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"github.com/nanadjei/go-jwt/initializers"
	"github.com/nanadjei/go-jwt/models"
)

func AuthCheck(context *gin.Context){
	// Get the token from the cookie
	tokenString, err := context.Cookie("Authorization")
	if err != nil {
		context.AbortWithStatus(http.StatusUnauthorized)
	}
	
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signup method %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			context.AbortWithStatus(http.StatusUnauthorized)
		}
		// Find the user with the token
		var user models.User
		initializers.DB.First(&user, claims["sub"]) //claims["exp"] represent the id of the user with the token
		if user.ID == 0 {
			context.AbortWithStatus(http.StatusUnauthorized)
		}
		// Attach the user
		context.Set("user", user)
		// Proceed with the request
		context.Next()
	} else {
		context.AbortWithStatus(http.StatusUnauthorized)
	}
	
}