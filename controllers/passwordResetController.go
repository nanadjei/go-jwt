package controllers

import (
	"text/template"

	"github.com/gin-gonic/gin"

	"github.com/nanadjei/go-jwt/response"
	"github.com/nanadjei/go-jwt/initializers"
	"github.com/nanadjei/go-jwt/models"
	"github.com/nanadjei/go-jwt/lib/mailer"
	"github.com/nanadjei/go-jwt/transformers"
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

	sendOTPcode(context, user.Email)
	
	response.Success(context, "Email Successfully sent.", transformers.UserTransform(user))
	return
}

func sendOTPcode(context *gin.Context, email string) {
	
	mailer := mailer.NewSMTPMail()
	// Parse the template
	t, _ := template.ParseFiles("/app/emails/passwordResetMail.html")

	// Data to be used in the template
	data := struct {
		Name    string
		Message string
	}{
		Name:    "Puneet Singh",
		Message: "This is a test message in an HTML template",
	}

	// Call the Send method on the mailer instance
	err := mailer.Send(email, "Password Reset Code", t, data)
	if err != nil {
		response.Error(context, "email", "The email was unable to send")
	}

}