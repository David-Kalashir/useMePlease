package repo

import (
	"context"
	"fmt"

	"github.com/David-Kalashir/crs-server/internal/entity"
	"github.com/David-Kalashir/crs-server/pkg/postgres"
)

const _defaultEntityCap = 64

type LoginRepo struct {
	*postgres.Postgres
}

func New(pg *postgres.Postgres) *LoginRepo {
	return &LoginRepo{pg}
}

func (r *LoginRepo) GetLogin(ctx context.Context) ([]entity.Login, error) {
	sql, _, err := r.Builder.
		Select("name, email").
		From("user").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("LoginRepo - GetLogin - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("LoginRepo - GetLogin - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	entities := make([]entity.Login, 0, _defaultEntityCap)

	for rows.Next() {
		e := entity.Login{}

		err = rows.Scan(&e.Email, &e.Name)
		if err != nil {
			return nil, fmt.Errorf("LoginRepo - GetLogin - rows.Scan: %w", err)
		}

		entities = append(entities, e)
	}

	return entities, nil
}
