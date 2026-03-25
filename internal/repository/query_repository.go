package repository

import (
	"context"
	"pg-query-to-email/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type QueryRepository struct {
	db *pgxpool.Pool
}

func NewQueryRepository(db *pgxpool.Pool) *QueryRepository {
	return &QueryRepository{db: db}
}

func (r *QueryRepository) GetUsers(ctx context.Context) ([]model.User, error) {
	rows, err := r.db.Query(ctx, GetUsersQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []model.User

	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.Nome, &user.Email, &user.Telefone, &user.CreatedAt); err != nil {
			return nil, err
		}
		results = append(results, user)
	}

	return results, nil
}
