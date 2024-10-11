package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Models:
type User struct {
	gorm.Model
	Username string `gorm:"unique" json:"username"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
}

var Database *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(sqlite.Open("./config/database/models.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	db.AutoMigrate(&User{}) // Include models as you create them

	Database = db // Moving the variable to the global scope
}

func UserExists(value, condition string) bool {
	var user User

	// Checks if user exists by email
	conditionString := condition + " = ?"
	result := Database.Where(conditionString, value).First(&user)

	return result.Error != gorm.ErrRecordNotFound // Returns true if finds a user, false if don't
}

func SearchUserByString(value, condition string) User {
	// Search user by any string type value

	var user User

	conditionString := condition + " = ?"
	Database.Where(conditionString, value).First(&user)
	return user
}

func SearchUserById(id uint) User {
	var user User

	Database.Where("id = ?", id).First(&user)
	return user
}
