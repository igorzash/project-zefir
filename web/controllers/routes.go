package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/igorzash/project-zefir/web/controllers/homecontroller"
	"github.com/igorzash/project-zefir/web/controllers/signincontroller"
	"github.com/igorzash/project-zefir/web/controllers/signupcontroller"
)

func SetUpRoutes(r *gin.Engine) {
	r.GET("/", homecontroller.Get)

	r.GET("/signin", signincontroller.Get)

	r.GET("/signup", signupcontroller.Get)
}
