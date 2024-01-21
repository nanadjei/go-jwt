package initializers

import (
	"gorm.io/driver/mysql"
  	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB
var err error

func DbConnect(){
	//"<db_user>:<db_password>@tcp(<db_host:localhost>:<db_port>)/<db_name>?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@tcp"+"("+os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT")+")/"+os.Getenv("DB_NAME")+"?charset=utf8mb4&parseTime=True&loc="+os.Getenv("DB_LOC")
  	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to db")
	}
}