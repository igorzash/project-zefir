package main

import (
	"log"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/igorzash/project-zefir/auth"
	"github.com/igorzash/project-zefir/db"
	"github.com/igorzash/project-zefir/user"
)

func main() {
	db.Connect()

	// Create a new Gin router
	r := gin.Default()

	authMiddleware := auth.GetMiddleware()

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
		r.GET("/user", user.HandleFind)
		r.GET("/user/:id", user.HandleGetByID)
		r.POST("/user", user.HandleCreate)
		r.PUT("/user/:id", user.HandleUpdate)
	}

	// Start the server
	r.Run("0.0.0.0:8080") // By default, this will run on localhost:8080
}
