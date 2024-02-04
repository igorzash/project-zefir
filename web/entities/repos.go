package entities

import (
	"github.com/igorzash/project-zefir/web/entities/followpkg"
	"github.com/igorzash/project-zefir/web/entities/userpkg"
)

type Repositories struct {
	UserRepo   userpkg.UserRepository
	FollowRepo followpkg.FollowRepository
}
