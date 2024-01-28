package middleware

import (
	"os"
	"net/http"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Guest(context *gin.Context){
	// Get the token from the cookie
	tokenString, err := context.Cookie("Authorization")
	if err != nil {
		// No token found, proceed with the request
		context.Next()
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected method %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check if the token has expired
		if time.Now().Unix() >= int64(claims["exp"].(float64)) {
			// Token has expired, proceed with the request
			context.Next()
			return
		}
	}

	// Token is invalid or not expired, abort the request
	context.AbortWithStatus(http.StatusUnauthorized)
	return
}