// database/models.go

package database

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string `json:"first_name"`
	Email     string `json:"email"`
}
