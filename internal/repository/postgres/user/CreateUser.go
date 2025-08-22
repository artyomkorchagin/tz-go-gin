package userpostgresql

import (
	"context"
	"fmt"

	"github.com/artyomkorchagin/tz-go-gin/internal/types"
)

func (r *Repository) CreateUser(ctx context.Context, u *types.User) error {
	query := `
		INSERT INTO users (login, full_name, gender, age, phone, email, avatar, is_active) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		`

	_, err := r.db.ExecContext(ctx, query,
		u.Login,
		u.FullName,
		u.Gender,
		u.Age,
		u.Phone,
		u.Email,
		u.Avatar,
		u.IsActive,
	)

	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}
