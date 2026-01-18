package users

import (
	"context"
	"errors"
	"fmt"
	"main/internal/domain"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetUserByEmail(ctx context.Context, pool *pgxpool.Pool, email string) (domain.User, error) {
	var user domain.User
	query := `SELECT id, email, password_hash FROM users WHERE email = $1`
	err := pool.QueryRow(ctx, query, email).Scan(
		&user.Id, &user.Email, &user.PasswordHash,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.User{}, nil
		}
		return domain.User{}, fmt.Errorf("failed to get user by email: %w", err)
	}

	return user, nil
}
