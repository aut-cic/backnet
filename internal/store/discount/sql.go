package discount

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/aut-cic/backnet/internal/model"
	"gorm.io/gorm"
)

type SQL struct {
	DB *gorm.DB
}

func NewSQL(db *gorm.DB) *SQL {
	return &SQL{
		DB: db,
	}
}

func (sql *SQL) List(ctx context.Context) ([]model.Discount, error) {
	var dfs []model.DiscountFactor

	if err := sql.DB.WithContext(ctx).Find(&dfs).Error; err != nil {
		return nil, fmt.Errorf("reading discounts from database failed %w", err)
	}

	var ds []model.Discount

	for _, discount := range dfs {
		switch discount.Type {
		case "date":
			date, err := time.Parse(time.DateOnly, discount.Value)
			if err != nil {
				return nil, fmt.Errorf("invalid date discount record %w", err)
			}

			ds = append(ds, model.DateDiscount{F: discount.Factor, Date: date})
		case "default":
			ds = append(ds, model.DefaultDiscount{F: discount.Factor})
		case "time":
			values := strings.Split(discount.Value, "-")

			since, err := time.Parse("15:04", values[0])
			if err != nil {
				return nil, fmt.Errorf("invalid time discount record %w", err)
			}

			until, err := time.Parse("15:04", values[1])
			if err != nil {
				return nil, fmt.Errorf("invalid time discount record %w", err)
			}

			ds = append(ds, model.TimeDiscount{F: discount.Factor, Since: since, Until: until})
		case "day_of_week":
			dw, err := strconv.Atoi(discount.Value)
			if err != nil {
				return nil, fmt.Errorf("invalid day_of_week discount record %w", err)
			}

			// nolint: gomnd
			ds = append(ds, model.DayDiscount{F: discount.Factor, Day: time.Weekday((dw + 1) % 7)})
		default:
			return nil, ErrInvalidDiscountType
		}
	}

	return ds, nil
}

func (sql *SQL) Create(ctx context.Context, m model.Discount) error {
	return nil
}
