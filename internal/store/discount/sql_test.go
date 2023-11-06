package discount_test

import (
	"context"
	"testing"
	"time"

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

func (suite *SQLDiscountSuite) TestListDiscounts() {
	require := suite.Require()

	discounts, err := suite.Store.List(context.Background())
	require.NoError(err)

	require.Len(discounts, 19)

	require.Equal("time", discounts[0].Type())
	require.Equal("23:00-23:59", discounts[0].Value())
	require.InEpsilon(0.4, discounts[0].Factor(), 0.001)

	require.Equal("day_of_week", discounts[1].Type())
	require.Equal("3", discounts[1].Value())
	require.InEpsilon(0.6, discounts[1].Factor(), 0.001)
	require.IsType(model.DayDiscount{Day: 0, F: 0}, discounts[1])
}

func (suite *SQLDiscountSuite) TestCreateDeleterDiscounts() {
	require := suite.Require()

	td := model.TimeDiscount{
		F:     1.0,
		Since: time.Now(),
		Until: time.Now().Add(time.Hour),
	}

	require.NoError(suite.Store.Create(context.Background(), td))
	require.NoError(suite.Store.Delete(context.Background(), td))
}

func TestSQLDiscountSuite(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(SQLDiscountSuite))
}
