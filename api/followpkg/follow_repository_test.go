package followpkg_test

import (
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/igorzash/project-zefir/followpkg"
	"github.com/igorzash/project-zefir/helpers"
	"github.com/igorzash/project-zefir/test"
	"github.com/igorzash/project-zefir/userpkg"
	"github.com/stretchr/testify/suite"
)

type FollowRepositorySuite struct {
	test.Suite
}

func TestFollow(t *testing.T) {
	suite.Run(t, new(FollowRepositorySuite))
}

func (suite *FollowRepositorySuite) TestFollows() {
	currentTime := time.Now().Format(time.RFC3339)
	users := [20]*userpkg.User{}
	for i := 0; i < 20; i++ {
		user := &userpkg.User{
			CreatedAt:    currentTime,
			UpdatedAt:    currentTime,
			PasswordHash: "0",
		}
		gofakeit.Struct(user)

		_, err := suite.Repos.UserRepo.Insert(user)
		suite.NoError(err)
		suite.NotNil(user.ID)

		users[i] = user
	}

	followMap := map[int]([]int){
		0:  {1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		1:  {2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
		2:  {3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
		3:  {4, 5, 6, 7, 8, 9, 10, 11, 12, 13},
		4:  {5, 6, 7, 8, 9, 10, 11, 12, 13, 14},
		5:  {6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
		6:  {7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
		7:  {8, 9, 10, 11, 12, 13, 14, 15, 16, 17},
		8:  {9, 10, 11, 12, 13, 14, 15, 16, 17, 18},
		9:  {10, 11, 12, 13, 14, 15, 16, 17, 18, 19},
		10: {11, 12, 13, 14, 15, 16, 17, 18, 19},
		11: {12, 13, 14, 15, 16, 17, 18, 19},
		12: {13, 14, 15, 16, 17, 18, 19},
		13: {14, 15, 16, 17, 18, 19},
		14: {15, 16, 17, 18, 19},
		15: {16, 17, 18, 19},
		16: {17, 18, 19},
		17: {18, 19},
		18: {19},
		19: {},
	}

	for followerID, followeeIDs := range followMap {
		for _, followeeID := range followeeIDs {
			follow := &followpkg.Follow{
				FollowerID: followerID,
				FolloweeID: followeeID,
				CreatedAt:  currentTime,
				UpdatedAt:  currentTime,
			}
			_, err := suite.Repos.FollowRepo.Insert(follow)
			suite.NoError(err)
		}
	}

	for followerID, followeeIDs := range followMap {
		usersFollowedBy, err := suite.Repos.FollowRepo.GetUsersFollowedBy(followerID, len(users), 0)
		suite.NoError(err)
		suite.Len(usersFollowedBy, len(followMap[followerID]))

		for _, userFollowedBy := range usersFollowedBy {
			suite.Contains(followeeIDs, userFollowedBy.ID)
		}

		usersFollowing, err := suite.Repos.FollowRepo.GetUserFollowers(followerID, len(users), 0)
		suite.NoError(err)

		for _, userFollowing := range usersFollowing {
			suite.Contains(followMap[userFollowing.ID], followerID)
		}

		for userID := 0; userID < len(followMap); userID++ {
			if helpers.Contains(followeeIDs, userID) {
				follow, err := suite.Repos.FollowRepo.GetByUsersIDs(followerID, userID)
				suite.NoError(err)
				suite.NotNil(follow)
				suite.Equal(followerID, follow.FollowerID)
				suite.Equal(userID, follow.FolloweeID)

				isMutual := helpers.Contains(followMap[userID], followerID)

				followState, err := suite.Repos.FollowRepo.GetFollowState(followerID, userID)
				suite.NoError(err)

				if isMutual {
					suite.Equal(followpkg.Mutual, followState)
				} else {
					suite.Equal(followpkg.Following, followState)
				}

				if isMutual {
					followState, err := suite.Repos.FollowRepo.GetFollowState(userID, followerID)
					suite.NoError(err)
					suite.Equal(followpkg.Mutual, followState)
				}
			} else {
				follow, err := suite.Repos.FollowRepo.GetByUsersIDs(followerID, userID)
				suite.NoError(err)
				suite.Nil(follow)

				followState, err := suite.Repos.FollowRepo.GetFollowState(followerID, userID)
				suite.NoError(err)
				suite.Equal(followpkg.NotFollowing, followState)
			}
		}
	}
}
