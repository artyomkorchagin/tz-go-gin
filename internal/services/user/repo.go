package userservice

import (
	"context"

	"github.com/artyomkorchagin/tz-go-gin/internal/types"
	"github.com/google/uuid"
)

type Reader interface {
	ReadUser(ctx context.Context, id uuid.UUID) (*types.User, error)
}

type Writer interface {
	CreateUser(ctx context.Context, user *types.User) error
}

type ReadWriter interface {
	Reader
	Writer
}
