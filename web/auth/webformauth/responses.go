package webformauth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func LoginResponse(c *gin.Context, code int, token string, expire time.Time) {
	c.Redirect(http.StatusMovedPermanently, "/home")
}

func UnauthorizedResponse(c *gin.Context, code int, message string) {
	c.Redirect(http.StatusMovedPermanently, "/signin?error=Incorrect email or password")
}
