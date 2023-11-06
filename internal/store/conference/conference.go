package conference

import (
	"context"
	"math/rand"

	"github.com/aut-cic/backnet/internal/model"
)

const PasswordLength = 7

type Conference interface {
	Create(ctx context.Context, name string, count int, prefix string) ([]model.Check, error)
	Delete(ctx context.Context, name string) error
	List(ctx context.Context, name string) ([]model.Check, error)
}

func RandomString(n int) string {
	letters := []rune("123456789")

	s := make([]rune, n)
	for i := range s {
		// nolint: gosec
		s[i] = letters[rand.Intn(len(letters))]
	}

	return string(s)
}
