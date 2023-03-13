package discount

import (
	"context"

	"github.com/aut-cic/backnet/internal/model"
)

type Discount interface {
	Create(context.Context, model.Discount) error
	Delete(context.Context, model.Discount) error
	List(context.Context) ([]model.Discount, error)
}
