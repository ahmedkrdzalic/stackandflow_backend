package database

import (
	"github.com/ahmedkrdzalic/StackAndFlow/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"log"
	"os"
)

var DB *gorm.DB

func Connect() {
	dns := os.Getenv("DATABASE_URL")
	conn, err := gorm.Open(mysql.Open(dns), &gorm.Config{})

	log.Printf("DATABASE is: " + dns)

	if err != nil {
		panic("Not connected to the database.\nERROR: ----- " + err.Error() + " -----")
	}

	DB = conn

	conn.AutoMigrate(&models.User{})
	conn.AutoMigrate(&models.Question{})
	conn.AutoMigrate(&models.Answer{})

}
