package database

import (
	"github.com/ahmedkrdzalic/StackAndFlow/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	conn, err := gorm.Open(mysql.Open("bfe61f88e37721:eada3ec7@us-cdbr-east-04.cleardb.com"), &gorm.Config{})

	if err != nil {
		panic("Not connected to the database.\nERROR: ----- " + err.Error() + " -----")
	}

	DB = conn

	conn.AutoMigrate(&models.User{})
	conn.AutoMigrate(&models.Question{})
	conn.AutoMigrate(&models.Answer{})

}
