package conference_test

import (
	"context"
	"testing"

	"github.com/aut-cic/backnet/internal/config"
	"github.com/aut-cic/backnet/internal/db"
	"github.com/aut-cic/backnet/internal/model"
	"github.com/aut-cic/backnet/internal/store/conference"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type CommonConferenceSuite struct{}

type SQLConferenceSuite struct {
	suite.Suite
	DB    *gorm.DB
	Store conference.Conference
}

func (suite *SQLConferenceSuite) SetupSuite() {
	cfg := config.New()

	db, err := db.New(cfg.Database)
	suite.Require().NoError(err)

	suite.DB = db
	suite.Store = conference.NewSQL(db)
}

func (suite *SQLConferenceSuite) TearDownSuite() {
}

func (suite *SQLConferenceSuite) TearDownTest() {
	suite.DB.Delete(new(model.Check))
	suite.DB.Delete(new(model.UserGroup))
}

func (suite *SQLConferenceSuite) TestNewConference() {
	require := suite.Require()

	users, err := suite.Store.Create(context.Background(), "elahe", 10, "Faculty")
	require.NoError(err)

	require.Len(users, 10)

	require.Equal("elahe00", users[0].Username)

	ug := new(model.UserGroup)

	require.NoError(suite.DB.Where("username = ?", "elahe09").Find(ug).Error)

	require.Equal("Faculty", ug.Groupname)
}

func (suite *SQLConferenceSuite) TestListConference() {
	require := suite.Require()

	_, err := suite.Store.Create(context.Background(), "parham", 10, "Faculty")
	require.NoError(err)

	suite.DB.Create(model.UserGroup{
		Username:  "parham.alvani",
		Groupname: "Faculty",
		ID:        0,
	})

	users, err := suite.Store.List(context.Background(), "parham")
	require.NoError(err)

	suite.Len(users, 10)
	suite.NotContains(users, "parham.alvani")
}

func (suite *SQLConferenceSuite) TestDeleteConference() {
	require := suite.Require()

	_, err := suite.Store.Create(context.Background(), "parham", 10, "Faculty")
	require.NoError(err)

	suite.DB.Create(model.UserGroup{
		Username:  "parham.alvani",
		Groupname: "Faculty",
		ID:        0,
	})

	require.NoError(suite.Store.Delete(context.Background(), "parham"))

	ug := new(model.UserGroup)
	require.NoError(suite.DB.Where("username = parham.alvani").Find(ug).Error)

	require.Equal("parham.alvani", ug.Username)
	require.Equal("Faculty", ug.Groupname)
}

func TestSQLConferenceSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(SQLConferenceSuite))
}
