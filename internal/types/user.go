package types

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserID           uuid.UUID `json:"user_id"`
	Login            string    `json:"login"`
	FullName         string    `json:"full_name"`
	Gender           string    `json:"gender"`
	Age              int       `json:"age"`
	Phone            string    `json:"phone"`
	Email            string    `json:"email"`
	Avatar           string    `json:"avatar"`
	RegistrationDate time.Time `json:"registration_date"`
	IsActive         bool      `json:"is_active"`
}
