package controllers

import (
	"net/http"
	"time"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v5"

	"github.com/nanadjei/go-jwt/models"
	"github.com/nanadjei/go-jwt/initializers"
	"github.com/nanadjei/go-jwt/response"
	"github.com/nanadjei/go-jwt/transformers"
)

func Signup(context *gin.Context){
	// Get the request body. ie: email, password
	var Body struct {
		Email string
		Password string
	}

	if context.Bind(&Body) != nil {
		response.Error(context, nil, "Failed to read body")
		return
	}
	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(Body.Password), 14)
	
	if err != nil {
		response.Error(context, "password", "Could not hash password")
		return 
	}
	// Store the user data in the db
	user := models.User{Email: Body.Email, Password: string(hash)}
	results := initializers.DB.Create(&user)
	
	if results.Error != nil {
		response.Error(context, "password", "Could not save data into database")
		return 
	}
	// Respond back to the request
	response.Success(context, "A new record was created successfully", transformers.UserTransform(user))
	return
}

func Signin(context *gin.Context) {
	var Body struct {
		Email string
		Password string
	}

	if context.Bind(&Body) != nil {
		response.Error(context, nil, "You failed to fill all required fields")
		return
	}

	var user models.User // The User Struct
	
	// Get email and password from request
	initializers.DB.First(&user, "email = ?", Body.Email)
	// Look out for the use with that credentials
	if user.ID == 0 {
		response.Error(context, "email", "There is no user associated with this email")
		return 
	}
	
	// Compare the request password with the user's passwords
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(Body.Password))

	if err != nil {
		response.Error(context, "password", "The password provided does not match the email provided")
		return
	}
	
	// Generate a JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,  jwt.MapClaims{
		"sub": user.ID,
		// Set expiration date
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	accessToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		response.Error(context, nil, "Access token failed to retreive")
		return
	}

	context.SetSameSite(http.SameSiteLaxMode)
	context.SetCookie("Authorization", accessToken, 3600*24*30, "", "", false, true)

	// Return JWT token as response
	data := make(map[string]interface{})
	data["access_token"] = accessToken
	data["user"] = user

	response.Success(context, "User successfully authenticated", data)
	return
}

func AuthUser(context *gin.Context) {
	authUser, _ := context.Get("user")
	context.JSON(http.StatusOK, gin.H{
		"authUser": authUser,
		"message": "You are already logged in",
	})
}

func Signout(context *gin.Context) {
	// Get token from Cookie
	_, err := context.Cookie("Authorization")
	if err != nil {
		context.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	context.SetCookie("Authorization", "", -1, "/", "", false, true)
	context.AbortWithStatus(http.StatusUnauthorized)
	return
}