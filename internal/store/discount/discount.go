package discount

import (
	"context"
	"errors"

	"github.com/aut-cic/backnet/internal/model"
)

var ErrInvalidDiscountType = errors.New("unknown discount type")

type Discount interface {
	Create(context.Context, model.Discount) error
	List(context.Context) ([]model.Discount, error)
}
