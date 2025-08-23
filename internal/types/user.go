package types

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	// Unique identifier of the user
	// swagger:ignore
	UserID uuid.UUID `json:"user_id,omitempty"`

	// Login/username of the user (required)
	// Required: true
	// minLength: 1
	// maxLength: 50
	Login string `json:"login" binding:"required"`

	// Full name of the user (required)
	// Required: true
	// minLength: 1
	// maxLength: 100
	FullName string `json:"full_name" binding:"required"`

	// Gender of the user
	// Enum: male,female,other
	Gender string `json:"gender,omitempty"`

	// Age of the user
	// Minimum: 0
	// Maximum: 150
	Age int `json:"age,omitempty"`

	// Phone number of the user
	// maxLength: 20
	Phone string `json:"phone,omitempty"`

	// Email address of the user
	// format: email
	// maxLength: 100
	Email string `json:"email,omitempty"`

	// URL to user's avatar image
	Avatar string `json:"avatar,omitempty"`

	// Registration date of the user
	// swagger:ignore
	RegistrationDate time.Time `json:"registration_date,omitempty"`

	// Whether the user account is active
	IsActive bool `json:"is_active,omitempty"`
}
