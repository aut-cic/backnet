package discount_test

import (
	"context"
	"testing"

	"github.com/aut-cic/backnet/internal/config"
	"github.com/aut-cic/backnet/internal/db"
	"github.com/aut-cic/backnet/internal/model"
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

	suite.Equal("time", discounts[0].Type())
	suite.Equal("23:00-23:59", discounts[0].Value())
	suite.Equal(0.4, discounts[0].Factor())

	suite.Equal("day_of_week", discounts[1].Type())
	suite.Equal("3", discounts[1].Value())
	suite.Equal(0.6, discounts[1].Factor())
	suite.IsType(model.DayDiscount{Day: 0, F: 0}, discounts[1])
}

func TestSQLDiscountSuite(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(SQLDiscountSuite))
}
