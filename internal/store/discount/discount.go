package discount

import (
	"context"

	"github.com/aut-cic/backnet/internal/model"
)

type Discount interface {
	Create(ctx context.Context, discont model.Discount) error
	Delete(ctx context.Context, discont model.Discount) error
	List(ctx context.Context) ([]model.Discount, error)
}
