package webformauth

import (
	"fmt"
	"os"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/igorzash/project-zefir/web/auth/userauth"
	"github.com/igorzash/project-zefir/web/entities"
)

func NewMiddleware(repos *entities.Repositories) (*jwt.GinJWTMiddleware, error) {
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
		IdentityKey:     userauth.IdentityKey,
		PayloadFunc:     userauth.PayloadFunc,
		IdentityHandler: userauth.IdentityHandler,
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
