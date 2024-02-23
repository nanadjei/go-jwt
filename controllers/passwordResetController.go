package controllers

import (
	"os"
	"text/template"
	"strconv"
	"time"
	// "fmt"

	"github.com/gin-gonic/gin"

	"github.com/nanadjei/go-jwt/response"
	"github.com/nanadjei/go-jwt/initializers"
	"github.com/nanadjei/go-jwt/models"
	"github.com/nanadjei/go-jwt/lib/mailer"
	"github.com/nanadjei/go-jwt/transformers"
	"github.com/nanadjei/go-jwt/helpers"
)

// Find the email of the user who want's to verify the email
func ForgotPassword(context *gin.Context){
	var Body struct {
		Email string
	}

	if context.Bind(&Body) != nil {
		response.Error(context, nil, "You failed to fill all required fields")
		return
	}

	var user models.User // The User Object
	initializers.DB.First(&user, "email = ?", Body.Email)

	if user.ID == 0 {
		response.Error(context, "user", "The user with this email was not found")
		return
	}

	// Generate OTP code
	sixDigitCode := helpers.GenerateOTPcode()

	// println("SIX DIGIT CODE:", sixDigitCode)

	// Convert the integer to string
	hashedCode, err := helpers.Encrypt(strconv.Itoa(sixDigitCode))
	
	if err != nil { println("Could not Encrypt") }

	// convert the env to int (Expiration Duration)
	envToInt, err := strconv.Atoi(os.Getenv("PASSWORD_RESET_TTL"))
	// Check for errors
	if err != nil { panic("Could not set env successfully") }

	// Multiply the int by seconds and get it in minutes. Eg: 900 * seconds = 15mins
	expiration :=  time.Duration(envToInt) * time.Second

	// Store the hashed 6 digits code for later verification
	StoreHashToRedis(context , hashedCode, expiration)

	go SendOTPcode(context, Body.Email, sixDigitCode, expiration)
	
	response.Success(context, "Email Successfully sent.", transformers.UserTransform(user))
	return
}

func SendOTPcode(context *gin.Context, email string, sixDigitCode int, expiration time.Duration) {
	
	mailer := mailer.NewSMTPMail()
	// Parse the template
	t, err := template.ParseFiles("/app/emails/passwordResetMail.html")
	if err != nil {
		println("Error parsing template:", err)
		return
	}

	// Data to be used in the template
	data := struct {
		AppName  string
		Email string
		Code int
		TTL time.Duration
	}{
		AppName: os.Getenv("APP_NAME"),
		Email: email,
		Code: sixDigitCode,
		TTL: expiration,
	}

	// Call the Send method on the mailer instance
	err = mailer.Send(email, "Password Reset Code", t, data)
	if err != nil {
		println("Error sending email:", err)
	}
}

func StoreHashToRedis(context *gin.Context, hashedString string, expiration time.Duration){
	err :=  initializers.Redis().Set(context, "email", hashedString, expiration).Err()
		if err != nil {
		response.Error(context, "redis", "Could not set in redis")
		return
	}
	return
}

func VerifyOTPcode(context *gin.Context) {
	var InputFields struct {
		Email string `json:"email" validate:"required,email"`
		Code string `json:"code" validate:"required"`
	}

	// Bind the input fields in the request with that of the InputFields struct
	if err := context.ShouldBind(&InputFields); err != nil {
		response.Error(context, "form", "Could not bind to input")
		return
	}

	// Since InputFields is of type var and has been bind, it will now carry the key, value pairs 
	// of the data sent from the request and can now be passed as an arg to the "ValidateInputs" method for validation
	if ok, errors := response.ValidateInputs(InputFields); !ok {
		response.ValidationError(context, errors, "Validation fails", 401)
		return
	}

	hashedCode, err := initializers.Redis().Get(context, "email").Result()

	if err != nil {
		response.Error(context, "code", "The code has expired.")
		return
	}

	val, err := helpers.Decrypt(hashedCode)
	code := InputFields.Code // The value of code passed from the request 

	switch val == code {
		case false:
		response.Error(context, "code", "The code does not match")
	case true:
		response.Success(context, "Code successfully matched", nil)
	}
}