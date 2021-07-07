package database

import (
	"github.com/ahmedkrdzalic/StackAndFlow/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"os"
)

var DB *gorm.DB

func Connect() {
	db_var := os.Getenv("CLEARDB_DATABASE_URL")
	conn, err := gorm.Open(mysql.Open(db_var), &gorm.Config{})

	if err != nil {
		panic("Not connected to the database.\nERROR: ----- " + err.Error() + " -----")
	}

	DB = conn

	conn.AutoMigrate(&models.User{})
	conn.AutoMigrate(&models.Question{})
	conn.AutoMigrate(&models.Answer{})

}
