package userservice

import (
	"context"
	"fmt"

	"github.com/artyomkorchagin/tz-go-gin/internal/types"
	"github.com/google/uuid"
)

func (s *Service) ReadUser(ctx context.Context, id string) (*types.User, error) {
	converted, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("error parsing string to UUID: %w", err)
	}
	return s.repo.ReadUser(ctx, converted)
}

// CreateUser creates a new user
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
func (s *Service) CreateUser(ctx context.Context, user *types.User) error {
	if err := s.repo.CreateUser(ctx, user); err != nil {
		return fmt.Errorf("error creating user: %w", err)
	}
	return nil
}
