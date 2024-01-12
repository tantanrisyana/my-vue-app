// database/init.go

package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:@tcp(127.0.0.1:3306)/userdb?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to the database")
		panic(err)
	}

	DB = db
	fmt.Println("Connected to the database")

	// AutoMigrate will create the table if it does not exist
	DB.AutoMigrate(&User{}) // Replace User with your model struct
}
