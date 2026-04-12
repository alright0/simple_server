package roles

import (
	"context"
	"fmt"
	"main/internal/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetRoles(ctx context.Context, pool *pgxpool.Pool) ([]domain.Role, error) {
	query := `SELECT roles.id, roles.name, roles.title FROM roles`

	rows, err := pool.Query(ctx, query)

	if err != nil {
		return nil, fmt.Errorf("failed to get roles: %w", err)
	}

	var roles []domain.Role
	for rows.Next() {
		var role domain.Role
		err := rows.Scan(&role.Id, &role.Name, &role.Title)
		if err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}

	return roles, nil
}
