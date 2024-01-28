package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Error response
// Parameters:
//	-	(context): The context of the request presented by gin
//	-	(code): Either a string or a nil value. Eg; 'email' or nill
//	-	(message): The message to respond with
func Error(context *gin.Context, code any, message string) {
	context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"code": code,
			"message": message,
		})
		return
}