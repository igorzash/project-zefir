package sign_up_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Post(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/")
}
