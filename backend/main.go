// backend/main.go

package main

import (
	"my-vue-app/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID        uint   `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
}

func main() {
	database.InitDB()

	r := gin.Default()

	r.GET("/api/users", func(c *gin.Context) {
		var users []User
		// Fetch only the desired columns
		database.DB.Select("id, email, first_name").Find(&users)
		c.JSON(http.StatusOK, users)
	})

	r.Run(":8080")
}
