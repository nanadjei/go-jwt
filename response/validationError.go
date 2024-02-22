package response

import (
	"strings"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Error response
// Parameters:
//	-	(context): The context of the request presented by gin
//	-	(message): The message to respond with
//	-	(responseData): The data returned from the server

func ValidationError(context *gin.Context, errorFields interface{},  message string, httpErrorCode int) {
	context.JSON(httpErrorCode | 401, gin.H{
			"success": false,
			"errors": errorFields,
			"message": message,
		})
		return 
	}

func ValidateInputs(dataset interface{}) (bool, map[string]string) {
	var validate *validator.Validate
	validate = validator.New()

	if err := validate.Struct(dataset); err != nil {
		// Validateion syntax is invalid
		if err, ok := err.(*validator.InvalidValidationError); ok {
			panic(err)
		}
		errors := make(map[string]string)
		reflected := reflect.ValueOf(dataset)

		for _, err := range err.(validator.ValidationErrors) {
			field, _ := reflected.Type().FieldByName(err.StructField())
			var name string
			if name = field.Tag.Get("json"); name == "" {
				name = strings.ToLower(err.StructField())
			}
			
			switch err.Tag() {
				case "required":
					errors[name] = "The " + name + " is required"
				break
				case "email":
					errors[name] = "The " + name + " should be a valid email"
				break
				case "eqField":
					errors[name] = "The " + name + " should be equal to the " + err.Param()
				break
				default: 
					errors[name] = "The " + name + " is invalid"
				break
			}
		}
		return false, errors
	}
	return true, nil
}