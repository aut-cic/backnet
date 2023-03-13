package discount

import (
	"context"
	"fmt"

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
		d, err := model.ToDiscount(discount)
		if err != nil {
			return nil, fmt.Errorf("casting discount database record failed %w", err)
		}

		ds = append(ds, d)
	}

	return ds, nil
}

func (sql *SQL) Create(ctx context.Context, m model.Discount) error {
	return nil
}
