package webformauth

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/igorzash/project-zefir/web/entities"
	"golang.org/x/crypto/bcrypt"
)

func Authenticator(repos *entities.Repositories) func(*gin.Context) (interface{}, error) {
	return func(c *gin.Context) (interface{}, error) {
		email := c.PostForm("email")
		password := c.PostForm("password")

		if email == "" || password == "" {
			return nil, jwt.ErrMissingLoginValues
		}

		user, _ := repos.UserRepo.GetByEmail(email)

		if user != nil {
			err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
			if err == nil {
				return user, nil
			}
		}

		return nil, jwt.ErrFailedAuthentication
	}
}
