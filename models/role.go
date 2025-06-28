package models

import "github.com/google/uuid"

type Role struct {
	ID       uuid.UUID `db:"id" json:"id"`
	RoleName string    `db:"role" json:"role"`
}
