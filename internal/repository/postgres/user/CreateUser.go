package userpostgresql

import (
	"context"
	"fmt"

	"github.com/artyomkorchagin/tz-go-gin/internal/types"
)

// CreateUser creates a new user in the database
// @Summary Create a new user
// @Description Creates a new user with the provided information
// @Tags users
// @Accept json
// @Produce json
// @Param user body types.User true "User information"
// @Success 200 {object} types.User "User created successfully"
// @Failure 400 {object} map[string]string "Bad request - invalid user data"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /users [post]
func (r *Repository) CreateUser(ctx context.Context, u *types.User) error {

	if u == nil {
		return types.ErrBadRequest(fmt.Errorf("user cant be nil"))
	}

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
		return types.ErrInternalServerError(fmt.Errorf("failed to create user: %w", err))
	}

	return nil
}
