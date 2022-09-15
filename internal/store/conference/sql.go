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

	tran := sql.DB.Begin()

	for i := 0; i < count; i++ {
		users[i] = model.Check{
			ID:        0,
			Username:  fmt.Sprintf("%s_%d", name, i),
			Attribute: "Cleartext-Password",
			Op:        ":=",
			Value:     RandomString(PasswordLength),
		}

		tran.WithContext(ctx).Create(&users[i])

		tran.WithContext(ctx).Create(&model.UserGroup{
			ID:        0,
			Username:  fmt.Sprintf("%s_%d", name, i),
			Groupname: groupName,
		})
	}

	if err := tran.Commit().Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (sql *SQL) Delete(ctx context.Context, name string) error {
	tran := sql.DB.Begin()

	tran.
		WithContext(ctx).
		Where("username LIKE ?", fmt.Sprintf("%s_%%", name)).
		Delete(new(model.Check))

	tran.
		WithContext(ctx).
		Where("username LIKE ?", fmt.Sprintf("%s_%%", name)).
		Delete(new(model.UserGroup))

	return tran.Commit().Error
}
