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

func (suite *SQLConferenceSuite) TestNewConference() {
	require := suite.Require()

	users, err := suite.Store.Create(context.Background(), "elahe", 10, "Faculty")
	require.NoError(err)

	require.Equal(10, len(users))

	require.Equal("elahe00", users[0].Username)

	ug := new(model.UserGroup)

	require.NoError(suite.DB.Where("username = ?", "elahe09").Find(ug).Error)

	require.Equal("Faculty", ug.Groupname)
}

func TestSQLConferenceSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(SQLConferenceSuite))
}
