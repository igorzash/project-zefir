package auth_web_form

import (
	"github.com/gin-gonic/gin"
)

func Authorizator() func(interface{}, *gin.Context) bool {
	return func(data interface{}, c *gin.Context) bool {
		return true
	}
}
