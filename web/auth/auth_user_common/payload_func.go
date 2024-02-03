package auth_user_common

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/igorzash/project-zefir/web/userpkg"
)

func PayloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(*userpkg.User); ok {
		return jwt.MapClaims{
			IdentityKey: v.Email,
		}
	}
	return jwt.MapClaims{}
}
