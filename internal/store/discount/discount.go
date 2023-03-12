package discount

import (
	"errors"

	"github.com/aut-cic/backnet/internal/model"
)

var ErrInvalidDiscountType = errors.New("unknown discount type")

type Discount interface {
	Create(model.Discount) error
	List() ([]model.Discount, error)
}
