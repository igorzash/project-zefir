package auth_user_common

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/igorzash/project-zefir/web/userpkg"
)

const IdentityKey = "email"

func IdentityHandler(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	return &userpkg.User{
		Email: claims[IdentityKey].(string),
	}
}
