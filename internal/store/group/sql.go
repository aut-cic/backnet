package group

import (
	"context"

	"github.com/aut-cic/backnet/internal/model"
	"gorm.io/gorm"
)

type SQL struct {
	DB     *gorm.DB
	groups []model.Package
}

func NewSQL(db *gorm.DB) *SQL {
	return &SQL{
		DB:     db,
		groups: nil,
	}
}

func (sql *SQL) List(ctx context.Context) ([]model.Package, error) {
	if sql.groups != nil {
		return sql.groups, nil
	}

	var groups []model.Package

	if err := sql.DB.WithContext(ctx).
		Find(&groups).Error; err != nil {
		return nil, err
	}

	sql.groups = groups

	return groups, nil
}
