package group

import (
	"context"

	"github.com/aut-cic/backnet/internal/model"
)

type Group interface {
	List(context.Context) ([]model.Package, error)
}
