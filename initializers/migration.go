package initializers

import "github.com/nanadjei/go-jwt/models"

func Migration(){
	var err error
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		panic("Error during migration: " + err.Error())
	}
}