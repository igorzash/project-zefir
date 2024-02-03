package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/igorzash/project-zefir/web/controllers/home_controller"
	"github.com/igorzash/project-zefir/web/controllers/sign_in_controller"
	"github.com/igorzash/project-zefir/web/controllers/sign_up_controller"
)

func SetUpRoutes(r *gin.Engine) {
	r.GET("/", home_controller.Get)

	r.GET("/signin", sign_in_controller.Get)

	r.GET("/signup", sign_up_controller.Get)
}
