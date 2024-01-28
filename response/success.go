package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Error response
// Parameters:
//	-	(context): The context of the request presented by gin
//	-	(message): The message to respond with
//	-	(responseData): The data returned from the server
func Success(context *gin.Context, message string, responseData map[string]interface{}) {
	context.JSON(http.StatusOK, gin.H{
			"success": true,
			"data": responseData,
			"message": message,
		})
		return 
	}