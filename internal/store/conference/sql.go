package conference

import (
	"context"
	"fmt"

	"github.com/aut-cic/backnet/internal/model"
	"github.com/pterm/pterm"
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

func (sql *SQL) Create(ctx context.Context, name string, count int, groupName string) ([]model.Check, error) {
	users := make([]model.Check, count)

	if err := sql.Delete(ctx, name); err != nil {
		pterm.Error.Printfln("cannot delete the old conference users %s", err)
	}

	tx := sql.DB.Begin()

	for i := 0; i < count; i++ {
		users[i] = model.Check{
			ID:        0,
			Username:  fmt.Sprintf("%s%02d", name, i),
			Attribute: "Cleartext-Password",
			Op:        ":=",
			Value:     RandomString(PasswordLength),
		}

		if err := tx.WithContext(ctx).Create(&users[i]).Error; err != nil {
			tx.Rollback()

			return nil, err
		}

		if err := tx.WithContext(ctx).Create(&model.UserGroup{
			ID:        0,
			Username:  fmt.Sprintf("%s%02d", name, i),
			Groupname: groupName,
		}).Error; err != nil {
			tx.Rollback()

			return nil, err
		}
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (sql *SQL) Delete(ctx context.Context, name string) error {
	tx := sql.DB.Begin()

	if err := tx.
		WithContext(ctx).
		Where("username LIKE ?", fmt.Sprintf("%s__", name)).
		Delete(new(model.Check)).Error; err != nil {
		tx.Rollback()

		return err
	}

	if err := tx.
		WithContext(ctx).
		Where("username LIKE ?", fmt.Sprintf("%s__", name)).
		Delete(new(model.UserGroup)).Error; err != nil {
		tx.Rollback()

		return err
	}

	return tx.Commit().Error
}

func (sql *SQL) List(ctx context.Context, name string) ([]model.Check, error) {
	var users []model.Check

	if err := sql.DB.WithContext(ctx).
		Where("username LIKE ?", fmt.Sprintf("%s__", name)).
		Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
