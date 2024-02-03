package auth_web_form

import (
	"fmt"
	"os"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/igorzash/project-zefir/web/auth/auth_user_common"
	"github.com/igorzash/project-zefir/web/repos"
)

func NewMiddleware(repos *repos.Repositories) (*jwt.GinJWTMiddleware, error) {
	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		return nil, fmt.Errorf("SECRET_KEY is not set")
	}

	domain := os.Getenv("DOMAIN")
	useHTTPS := true
	if domain == "" {
		domain = "localhost"
		useHTTPS = false
	}

	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:           "web",
		Key:             []byte(secretKey),
		Timeout:         time.Hour,
		MaxRefresh:      time.Hour,
		IdentityKey:     auth_user_common.IdentityKey,
		PayloadFunc:     auth_user_common.PayloadFunc,
		IdentityHandler: auth_user_common.IdentityHandler,
		Authenticator:   Authenticator(repos),
		Authorizator:    Authorizator(),
		LoginResponse:   LoginResponse,
		Unauthorized:    UnauthorizedResponse,
		SendCookie:      true,
		CookieName:      "jwt_token",
		CookieDomain:    "localhost",
		CookieHTTPOnly:  true,
		SecureCookie:    useHTTPS,
		TimeFunc:        time.Now,
	})

	if err != nil {
		return nil, err
	}

	return authMiddleware, nil
}
