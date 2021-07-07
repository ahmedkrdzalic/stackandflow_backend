package database

import (
	"github.com/ahmedkrdzalic/StackAndFlow/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	conn, err := gorm.Open(mysql.Open("root:@/MOP_DB?parseTime=true"), &gorm.Config{})

	if err != nil {
		panic("Not connected to the database.\nERROR: ----- " + err.Error() + " -----")
	}

	DB = conn

	conn.AutoMigrate(&models.User{})
	conn.AutoMigrate(&models.Question{})
	conn.AutoMigrate(&models.Answer{})

}
