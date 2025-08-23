package userpostgresql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/artyomkorchagin/tz-go-gin/internal/types"
	"github.com/google/uuid"
)

// ReadUser retrieves a user by ID from the database
// @Summary Get user by ID
// @Description Retrieves detailed information about a user by their unique identifier
// @Tags users
// @Accept json
// @Produce json
// @Param user_id path string true "User UUID" Format(uuid)
// @Success 200 {object} types.User "User found successfully"
// @Failure 400 {object} map[string]string "Bad request - invalid UUID format"
// @Failure 404 {object} map[string]string "User not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /users/{user_id} [get]
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
