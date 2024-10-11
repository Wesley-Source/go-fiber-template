package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique" json:"username"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
}

var Database *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(sqlite.Open("./config/database/db.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	db.AutoMigrate(&User{}) // Include all the models

	Database = db
}
