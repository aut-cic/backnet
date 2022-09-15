package conference

import (
	"context"
	"math/rand"

	"github.com/aut-cic/backnet/internal/model"
)

const PasswordLength = 7

type Conference interface {
	Create(context.Context, string, int, string) ([]model.Check, error)
	Delete(context.Context, string) error
	List(context.Context, string) []model.Check
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
