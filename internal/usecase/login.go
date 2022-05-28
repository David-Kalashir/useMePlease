package usecase

import (
	"context"
	"fmt"

	"github.com/David-Kalashir/crs-server/internal/entity"
)

// TranslationUseCase -.
type TranslationUseCase struct {
	repo LoginRepo
}

// New -.
func New(r LoginRepo) *TranslationUseCase {
	return &TranslationUseCase{
		repo: r,
	}
}

// Login from store.
func (uc *TranslationUseCase) Login(ctx context.Context) ([]entity.Login, error) {
	translations, err := uc.repo.GetLogin(ctx)
	if err != nil {
		return nil, fmt.Errorf("LoginUseCase - Login - s.repo.Login: %w", err)
	}

	return translations, nil
}
