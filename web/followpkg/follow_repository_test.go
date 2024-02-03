package followpkg_test

import (
	"testing"

	"github.com/igorzash/project-zefir/web/followpkg"
	"github.com/igorzash/project-zefir/web/helpers"
	"github.com/igorzash/project-zefir/web/test"
	"github.com/igorzash/project-zefir/web/userpkg"
	"github.com/stretchr/testify/suite"
)

type FollowRepositorySuite struct {
	test.Suite
}

func TestFollow(t *testing.T) {
	suite.Run(t, new(FollowRepositorySuite))
}

const NumUsers = 100

func (suite *FollowRepositorySuite) TestFollows() {
	users := [NumUsers]*userpkg.User{}

	for i := 0; i < len(users); i++ {
		user := suite.NewUser()
		suite.App.Repos.UserRepo.Insert(user)
		users[i] = user
	}

	followMap := make(map[int][]int, NumUsers)
	for i := 0; i < NumUsers; i++ {
		for j := i + 1; j < NumUsers && j <= i+10; j++ {
			followMap[i] = append(followMap[i], j)
		}
	}

	for followerID, followeeIDs := range followMap {
		for _, followeeID := range followeeIDs {
			follow := followpkg.NewFollow(followerID, followeeID)
			_, err := suite.App.Repos.FollowRepo.Insert(follow)
			suite.NoError(err)
		}
	}

	for followerID, followeeIDs := range followMap {
		usersFollowedBy, err := suite.App.Repos.FollowRepo.GetUsersFollowedBy(followerID, len(users), 0)
		suite.NoError(err)
		suite.Len(usersFollowedBy, len(followMap[followerID]))

		for _, userFollowedBy := range usersFollowedBy {
			suite.Contains(followeeIDs, userFollowedBy.ID)
		}

		usersFollowing, err := suite.App.Repos.FollowRepo.GetUserFollowers(followerID, len(users), 0)
		suite.NoError(err)

		for _, userFollowing := range usersFollowing {
			suite.Contains(followMap[userFollowing.ID], followerID)
		}

		for userID := 0; userID < len(followMap); userID++ {
			if helpers.Contains(followeeIDs, userID) {
				follow, err := suite.App.Repos.FollowRepo.GetByUsersIDs(followerID, userID)
				suite.NoError(err)
				suite.NotNil(follow)
				suite.Equal(followerID, follow.FollowerID)
				suite.Equal(userID, follow.FolloweeID)

				isMutual := helpers.Contains(followMap[userID], followerID)

				followState, err := suite.App.Repos.FollowRepo.GetFollowState(followerID, userID)
				suite.NoError(err)

				if isMutual {
					suite.Equal(followpkg.Mutual, followState)
				} else {
					suite.Equal(followpkg.Following, followState)
				}

				if isMutual {
					followState, err := suite.App.Repos.FollowRepo.GetFollowState(userID, followerID)
					suite.NoError(err)
					suite.Equal(followpkg.Mutual, followState)
				}
			} else {
				follow, err := suite.App.Repos.FollowRepo.GetByUsersIDs(followerID, userID)
				suite.NoError(err)
				suite.Nil(follow)

				followState, err := suite.App.Repos.FollowRepo.GetFollowState(followerID, userID)
				suite.NoError(err)
				suite.Equal(followpkg.NotFollowing, followState)
			}
		}
	}
}

func (suite *FollowRepositorySuite) TestUnFollow() {
	users := [NumUsers]*userpkg.User{}

	for i := 0; i < len(users); i++ {
		user := suite.NewUser()
		suite.App.Repos.UserRepo.Insert(user)
		users[i] = user
	}

	for i := 0; i < len(users); i++ {
		for j := i; j < len(users); j++ {
			if i == j {
				continue
			}

			_, err := suite.App.Repos.FollowRepo.Delete(i, j)
			suite.NoError(err)

			follow := followpkg.NewFollow(i, j)
			_, err = suite.App.Repos.FollowRepo.Insert(follow)
			suite.NoError(err)

			_, err = suite.App.Repos.FollowRepo.Delete(i, j)
			suite.NoError(err)

			followState, err := suite.App.Repos.FollowRepo.GetFollowState(i, j)
			suite.NoError(err)
			suite.Equal(followpkg.NotFollowing, followState)
		}
	}
}
