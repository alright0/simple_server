package user_services

import (
	"context"
	"fmt"
	"main/internal/domain"
	"main/internal/dto"
	"main/internal/repositories/users"

	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

func LoginUser(ctx context.Context, pool *pgxpool.Pool, userData dto.LoginRequest) (domain.User, error) {
	user, err := users.GetUserByEmail(ctx, pool, userData.Email)
	if err != nil {
		return domain.User{}, fmt.Errorf("error while getting user: %w", err)
	}
	if user == (domain.User{}) {
		return domain.User{}, fmt.Errorf("invalid credentials")
	}
	if user.IsDeleted {
		return domain.User{}, fmt.Errorf("Users is inactive")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(userData.Password)); err != nil {
		return domain.User{}, fmt.Errorf("invalid credentials")
	}
	return user, nil
}
