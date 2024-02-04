package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/igorzash/project-zefir/web/auth/webformauth"
	"github.com/igorzash/project-zefir/web/entities"
)

func SetUpRoutes(r *gin.Engine, repos *entities.Repositories) error {
	authWebFormMiddleware, err := webformauth.NewMiddleware(repos)
	if err != nil {
		return err
	}

	err = authWebFormMiddleware.MiddlewareInit()
	if err != nil {
		return err
	}

	r.POST("/signin", authWebFormMiddleware.LoginHandler)
	r.GET("/refresh_token", authWebFormMiddleware.RefreshHandler)

	return nil
}
