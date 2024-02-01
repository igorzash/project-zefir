package userpkg_test

import (
	"testing"
	"time"

	"github.com/igorzash/project-zefir/test"
	"github.com/igorzash/project-zefir/userpkg"
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
	user, err := suite.Repos.UserRepo.GetByEmail(gofakeit.Email())
	suite.NoError(err)
	suite.Nil(user)
}

func (suite *UserRepositorySuite) TestGetAfterCreateAndUpdate() {
	currentTime := time.Now().Format(time.RFC3339)
	user := userpkg.User{
		CreatedAt:    currentTime,
		UpdatedAt:    currentTime,
		PasswordHash: "0",
	}
	gofakeit.Struct(&user)
	_, err := suite.Repos.UserRepo.Insert(&user)
	suite.NoError(err)
	suite.NotNil(user.ID)

	assertUserMatchesDbVersion := func() {
		userFromDb, err := suite.Repos.UserRepo.GetByEmail(user.Email)
		suite.NoError(err)
		suite.NotNil(userFromDb)
		suite.Equal(&user, userFromDb)

		userFromDb, err = suite.Repos.UserRepo.GetByID(user.ID)
		suite.NoError(err)
		suite.NotNil(userFromDb)
		suite.Equal(&user, userFromDb)
	}
	assertUserMatchesDbVersion()

	user.UpdatedAt = time.Now().Format(time.RFC3339)
	oldNickname := user.Nickname
	for oldNickname == user.Nickname {
		user.Nickname = gofakeit.Username()
	}
	_, err = suite.Repos.UserRepo.Update(&user)
	suite.NoError(err)
	assertUserMatchesDbVersion()
}
