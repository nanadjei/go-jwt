package transformers

import "github.com/nanadjei/go-jwt/models"

// Make a user transformer to spell out how the data should be presented
// Parameters:
//	-	user (string): The user model
// Return 
//	-	map[string]interface{}
func UserTransform(user models.User) map[string]interface{} {
	userMap := map[string]interface{}{
		"id": user.ID,
		"email": user.Email,
		"createdAt": user.CreatedAt,
	}
	return userMap
}