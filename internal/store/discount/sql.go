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

	ds := make([]model.Discount, len(dfs))

	for i, discount := range dfs {
		d, err := model.ToDiscount(discount)
		if err != nil {
			return nil, fmt.Errorf("casting discount database record failed %w", err)
		}

		ds[i] = d
	}

	return ds, nil
}

func (sql *SQL) Create(ctx context.Context, m model.Discount) error {
	if err := sql.DB.WithContext(ctx).Create(&model.DiscountFactor{
		ID:     0,
		Factor: m.Factor(),
		Value:  m.Value(),
		Type:   m.Type(),
	}).Error; err != nil {
		return fmt.Errorf("inserting discount into database failed %w", err)
	}

	return nil
}

func (sql *SQL) Delete(ctx context.Context, m model.Discount) error {
	if err := sql.DB.WithContext(ctx).
		Where("type = ? AND value = ?", m.Type(), m.Value()).
		Delete(new(model.DiscountFactor)).
		Error; err != nil {
		return fmt.Errorf("deleting discount from database failed %w", err)
	}

	return nil
}
