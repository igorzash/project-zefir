package main

import (
	"log"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type login struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type User struct {
	Id           int    `db:"id"`
	Email        string `db:"email" json:"-"`
	CreatedAt    string `db:"created_at" json:"createdAt"`
	UpdatedAt    string `db:"updated_at" json:"updatedAt"`
	Nickname     string `db:"nickname"`
	PasswordHash string `db:"password_hash" json:"-"`
}

func main() {
	ConnectToDb()

	// Create a new Gin router
	r := gin.Default()

	authMiddleware := SetUpAuthMiddleware()

	// When you use jwt.New(), the function is already automatically called for checking,
	// which means you don't need to call it again.
	errInit := authMiddleware.MiddlewareInit()

	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}

	r.POST("/login", authMiddleware.LoginHandler)

	r.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	auth := r.Group("/auth")
	// Refresh time can be longer than token timeout
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		r.GET("/users", HandleGetUsers)
		r.GET("/users/:id", HandleGetUserByID)
		r.POST("/users", HandleCreateUser)
		r.PUT("/users/:id", HandleUpdateUser)
	}

	// Start the server
	r.Run("0.0.0.0:8080") // By default, this will run on localhost:8080
}
