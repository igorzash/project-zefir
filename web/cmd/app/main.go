package main

import (
	"log"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/igorzash/project-zefir/web/auth"
	"github.com/igorzash/project-zefir/web/db"
	"github.com/igorzash/project-zefir/web/repos"
	"github.com/igorzash/project-zefir/web/userpkg"
)

func main() {
	dbConn := db.Connect()

	repos, err := repos.NewRepositories(dbConn)
	if err != nil {
		log.Fatal("Failed to initialize repositories: " + err.Error())
	}

	// Create a new Gin router
	r := gin.Default()

	authMiddleware := auth.GetMiddleware(repos)

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
		r.GET("/user", userpkg.HandleFind)
		r.GET("/user/:id", userpkg.HandleGetByID)
		r.POST("/user", userpkg.HandleCreate)
		r.PUT("/user/:id", userpkg.HandleUpdate)
	}

	// Start the server
	r.Run("0.0.0.0:8080") // By default, this will run on localhost:8080
}
