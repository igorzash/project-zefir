package test_test

import (
	"testing"

	"github.com/igorzash/project-zefir/web/test"
	"github.com/stretchr/testify/suite"
)

type TestTestSuite struct {
	test.Suite
}

func TestTest(t *testing.T) {
	suite.Run(t, new(TestTestSuite))
}

func (suite *TestTestSuite) TestAppInitializedOK() {
	suite.NotNil(suite.App)
	suite.NotNil(suite.App.Repos)
	suite.NotNil(suite.App.DBConn)
	suite.NotNil(suite.App.R)
}
