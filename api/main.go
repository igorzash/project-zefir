package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	ConnectToDb()

	// Create a new Gin router
	r := gin.Default()

	r.GET("/users", HandleGetUsers)
	r.GET("/users/:id", HandleGetUserByID)
	r.POST("/users", HandleCreateUser)
	r.PUT("/users/:id", HandleUpdateUser)

	// Start the server
	r.Run("0.0.0.0:8080") // By default, this will run on localhost:8080
}
