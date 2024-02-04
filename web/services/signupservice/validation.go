package signupservice

import (
	"fmt"

	"github.com/igorzash/project-zefir/web/entities/userpkg"
)

type SignupService interface {
	Signup(email string, nickname string, password string) error
}

type DefaultSignupService struct {
	// Add fields for dependencies here, like a database connection
}

func NewDefaultSignupService() SignupService {
	return &DefaultSignupService{}
}

func (s *DefaultSignupService) Signup(email string, nickname string, password string) error {
	_, err := userpkg.NewUser(email, nickname, password)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}
