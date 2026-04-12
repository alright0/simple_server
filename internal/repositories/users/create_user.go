package users

import (
	"context"
	"main/internal/dto"
	"main/internal/utils"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateUser(ctx context.Context, pool *pgxpool.Pool, user dto.CreateUserRequest) (int, error) {
	var userId int
	passwordHash, err := utils.HashPassword(user.Password)
	if err != nil {
		return 0, err
	}

	user.Email = strings.ToLower(strings.TrimSpace(user.Email))

	query := `INSERT INTO users (email, password_hash, role_id) 
			    VALUES ($1, $2, (SELECT id FROM roles where name='user')) 
			    RETURNING id`

	err = pool.QueryRow(ctx, query, user.Email, passwordHash).Scan(&userId)
	if err != nil {
		return 0, err
	}

	return userId, nil
}
