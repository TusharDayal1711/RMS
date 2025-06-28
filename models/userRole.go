package models

import (
	"github.com/google/uuid"
	"time"
)

type UserRole struct {
	ID        uuid.UUID  `db:"id" json:"id"`
	UserID    uuid.UUID  `db:"user_id" json:"user_id"`
	RoleID    uuid.UUID  `db:"role_id" json:"role_id"`
	CreatedBy *uuid.UUID `db:"created_by" json:"created_by,omitempty"` //can be null in case user created account themselves and will be omitted if null
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
}

type MultiRole struct {
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	Password string   `json:"password"`
	Roles    []string `json:"roles"`
}
