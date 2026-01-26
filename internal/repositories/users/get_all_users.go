package users

import (
	"context"
	"fmt"
	"main/internal/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetUsers(ctx context.Context, pool *pgxpool.Pool) ([]domain.UserForList, error) {
	query := `
		SELECT 
		      users.id
		    , users.email
		    , users.is_deleted
		    , users.updated_at
		    , users.created_at
		    , roles.name
		    , roles.title 
		FROM users
		JOIN roles ON roles.id = users.role_id
		`

	rows, err := pool.Query(ctx, query)

	if err != nil {
		return nil, fmt.Errorf("failed to get users: %w", err)
	}

	var users []domain.UserForList
	for rows.Next() {
		var user domain.UserForList
		err := rows.Scan(&user.Id, &user.Email, &user.IsDeleted, &user.UpdatedAt, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
