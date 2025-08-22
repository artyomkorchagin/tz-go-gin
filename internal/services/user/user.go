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

func (s *Service) CreateUser(ctx context.Context, user *types.User) error {
	if err := s.repo.CreateUser(ctx, user); err != nil {
		return fmt.Errorf("error creating user: %w", err)
	}
	return nil
}
