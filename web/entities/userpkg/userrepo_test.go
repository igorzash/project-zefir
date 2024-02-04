package userpkg_test

import (
	"testing"
	"time"

	"github.com/igorzash/project-zefir/web/entities/userpkg"
	"github.com/igorzash/project-zefir/web/test"
	"github.com/stretchr/testify/suite"

	"github.com/brianvoe/gofakeit/v6"
)

type UserRepositorySuite struct {
	test.Suite
}

func TestUser(t *testing.T) {
	suite.Run(t, new(UserRepositorySuite))
}

func (suite *UserRepositorySuite) TestGetByEmailNil() {
	user, err := suite.App.Repos.UserRepo.GetByEmail(gofakeit.Email())
	suite.NoError(err)
	suite.Nil(user)
}

func (suite *UserRepositorySuite) TestGetAfterCreateAndUpdate() {
	user, err := userpkg.NewUser(gofakeit.Email(), gofakeit.Username(), gofakeit.Password(true, true, true, false, false, 12))
	suite.NoError(err)

	_, err = suite.App.Repos.UserRepo.Insert(user)
	suite.NoError(err)
	suite.NotNil(user.ID)

	assertUserMatchesDbVersion := func() {
		userFromDb, err := suite.App.Repos.UserRepo.GetByEmail(user.Email)
		suite.NoError(err)
		suite.NotNil(userFromDb)
		suite.Equal(user, userFromDb)

		userFromDb, err = suite.App.Repos.UserRepo.GetByID(user.ID)
		suite.NoError(err)
		suite.NotNil(userFromDb)
		suite.Equal(user, userFromDb)
	}
	assertUserMatchesDbVersion()

	user.UpdatedAt = time.Now().Format(time.RFC3339)
	oldNickname := user.Nickname
	for oldNickname == user.Nickname {
		user.Nickname = gofakeit.Username()
	}
	_, err = suite.App.Repos.UserRepo.Update(user)
	suite.NoError(err)
	assertUserMatchesDbVersion()
}
