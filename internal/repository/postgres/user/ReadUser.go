package userpostgresql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/artyomkorchagin/tz-go-gin/internal/types"
	"github.com/google/uuid"
)

func (r *Repository) ReadUser(ctx context.Context, user_id uuid.UUID) (*types.User, error) {
	query := `
		SELECT user_id, login, full_name, gender, age, phone, email, avatar, registration_date, is_active 
		FROM users 
		WHERE user_id = $1`

	var user types.User
	err := r.db.QueryRowContext(ctx, query, user_id).Scan(
		&user.UserID,
		&user.Login,
		&user.FullName,
		&user.Gender,
		&user.Age,
		&user.Phone,
		&user.Email,
		&user.Avatar,
		&user.RegistrationDate,
		&user.IsActive,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found: %w", err)
		}
		return nil, fmt.Errorf("failed to read user: %w", err)
	}

	return &user, nil
}
