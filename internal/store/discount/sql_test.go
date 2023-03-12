package discount_test

import (
	"context"
	"testing"

	"github.com/aut-cic/backnet/internal/config"
	"github.com/aut-cic/backnet/internal/db"
	"github.com/aut-cic/backnet/internal/store/discount"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type CommonDiscountSuite struct{}

type SQLDiscountSuite struct {
	suite.Suite
	DB    *gorm.DB
	Store discount.Discount
}

func (suite *SQLDiscountSuite) SetupSuite() {
	cfg := config.New()

	db, err := db.New(cfg.Database)
	suite.Require().NoError(err)

	suite.DB = db
	suite.Store = discount.NewSQL(db)
}

func (suite *SQLDiscountSuite) TearDownSuite() {
}

func (suite *SQLDiscountSuite) AfterTest() {
}

func (suite *SQLDiscountSuite) TestListConference() {
	require := suite.Require()

	discounts, err := suite.Store.List(context.Background())
	require.NoError(err)

	suite.Len(discounts, 19)

	suite.Equal(discounts[0].Type(), "time")
	suite.Equal(discounts[0].Value(), "23:00-23:59")
	suite.Equal(discounts[0].Factor(), 0.4)
}

func TestSQLDiscountSuite(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(SQLDiscountSuite))
}
