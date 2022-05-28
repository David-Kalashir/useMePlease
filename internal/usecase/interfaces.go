// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"

	"github.com/David-Kalashir/crs-server/internal/entity"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase_test

type (
	// Translation -.
	Rigestery interface {
		Login(context.Context) ([]entity.Login, error)
	}

	// TranslationRepo -.
	LoginRepo interface {
		GetLogin(context.Context) ([]entity.Login, error)
	}
)
