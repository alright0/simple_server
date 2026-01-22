package users

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func SoftDeleteUser(ctx context.Context, pool *pgxpool.Pool, userId int) (bool, error) {
	query := `UPDATE users SET is_deleted=TRUE WHERE id=$1 RETURNING id`
	err := pool.QueryRow(ctx, query, userId).Scan(&userId)
	if err != nil {
		return false, err
	}

	return true, nil
}
