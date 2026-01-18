package users

import (
	"context"
	"main/internal/dto"
	"main/internal/utils"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateUser(ctx context.Context, pool *pgxpool.Pool, user dto.CreateUserRequest) (int, error) {
	var userId int
	passwordHash := utils.HashPassword(user.Password)

	query := `INSERT INTO users (email, password_hash) VALUES ($1, $2) RETURNING id`
	err := pool.QueryRow(ctx, query, user.Email, passwordHash).Scan(&userId)
	if err != nil {
		return 0, err
	}

	return userId, nil
}
