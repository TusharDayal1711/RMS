package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID         uuid.UUID  `db:"id" json:"id"`
	Name       string     `db:"name" json:"name"`
	Email      string     `db:"email" json:"email"`
	Password   string     `db:"password" json:"password"`
	CreatedAt  time.Time  `db:"created_at" json:"created_at"`
	ArchivedAt *time.Time `db:"archived_at" json:"archived_at,omitempty"`
}

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SubAdminReq struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserReq struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserCreatedBy struct {
	ID        uuid.UUID `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	Email     string    `db:"email" json:"email"`
	Role      string    `db:"role" json:"role,"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

type SubAdminCreatedBy struct {
	ID        uuid.UUID `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	Email     string    `db:"email" json:"email"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}
