package home_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
